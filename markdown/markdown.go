package markdown

import (
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/devallabharath/md-notes/utils"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
)

type mdRenderer []widget.RichTextSegment

func (m *mdRenderer) AddOptions(...renderer.Option) {}

func (m *mdRenderer) Render(_ io.Writer, source []byte, n ast.Node) error {
	segs, err := renderNode(source, n, false)
	*m = segs
	return err
}

func renderNode(source []byte, n ast.Node, blockquote bool) ([]widget.RichTextSegment, error) {
	switch t := n.(type) {
	case *ast.Document:
		return renderChildren(source, n, blockquote)
	case *ast.Paragraph:
		children, err := renderChildren(source, n, blockquote)
		if !blockquote {
			linebreak := &widget.TextSegment{Style: MdStyles.Paragraph}
			children = append(children, linebreak)
		}
		return children, err
	case *ast.List:
		items, err := renderChildren(source, n, blockquote)
		return []widget.RichTextSegment{
			&widget.ListSegment{Items: items, Ordered: t.Marker != '*' && t.Marker != '-' && t.Marker != '+'},
		}, err
	case *ast.ListItem:
		texts, err := renderChildren(source, n, blockquote)
		return []widget.RichTextSegment{&widget.ParagraphSegment{Texts: texts}}, err
	case *ast.TextBlock:
		return renderChildren(source, n, blockquote)
	case *ast.Heading:
		text := forceIntoHeadingText(source, n)
		switch t.Level {
		case 1:
			return []widget.RichTextSegment{&widget.TextSegment{Style: MdStyles.Heading1, Text: text}}, nil
		case 2:
			return []widget.RichTextSegment{&widget.TextSegment{Style: MdStyles.Heading2, Text: text}}, nil
		default:
			textSegment := widget.TextSegment{Style: MdStyles.Heading, Text: text}
			return []widget.RichTextSegment{&textSegment}, nil
		}
	case *ast.ThematicBreak:
		return []widget.RichTextSegment{
			&widget.TextSegment{Style: MdStyles.Emptyline, Text: "\n"},
			&widget.SeparatorSegment{},
			&widget.TextSegment{Style: MdStyles.Emptyline, Text: "\n"},
		}, nil
	case *ast.Link:
		link, _ := url.Parse(string(t.Destination))
		text := forceIntoText(source, n)
		return []widget.RichTextSegment{&widget.HyperlinkSegment{Alignment: fyne.TextAlignLeading, Text: text, URL: link}}, nil
	case *ast.CodeSpan:
		text := forceIntoText(source, n)
		return []widget.RichTextSegment{&widget.TextSegment{Style: widget.RichTextStyleCodeInline, Text: text}}, nil
	case *ast.CodeBlock, *ast.FencedCodeBlock:
		var data []byte
		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			data = append(data, line.Value(source)...)
		}
		if len(data) == 0 {
			return nil, nil
		}
		if data[len(data)-1] == '\n' {
			data = data[:len(data)-1]
		}
		return []widget.RichTextSegment{
			&widget.TextSegment{Style: MdStyles.Emptyline, Text: "\n"},
			&widget.TextSegment{Style: MdStyles.Codeblock, Text: string(data)},
		}, nil
	case *ast.Emphasis:
		text := string(forceIntoText(source, n))
		switch t.Level {
		case 2:
			return []widget.RichTextSegment{&widget.TextSegment{Style: widget.RichTextStyleStrong, Text: text}}, nil
		default:
			return []widget.RichTextSegment{&widget.TextSegment{Style: widget.RichTextStyleEmphasis, Text: text}}, nil
		}
	case *ast.Text:
		text := string(t.Text(source))
		if text == "" {
			// These empty text elements indicate single line breaks after non-text elements in goldmark.
			return []widget.RichTextSegment{&widget.TextSegment{Style: widget.RichTextStyleInline, Text: " "}}, nil
		}
		text = suffixSpaceIfAppropriate(text, n)
		if blockquote {
			return []widget.RichTextSegment{&widget.TextSegment{Style: widget.RichTextStyleBlockquote, Text: text}}, nil
		}
		return []widget.RichTextSegment{&widget.TextSegment{Style: widget.RichTextStyleInline, Text: text}}, nil
	case *ast.Blockquote:
		return renderChildren(source, n, true)
	case *ast.Image:
		dest := string(t.Destination)
		var url fyne.URI
		if filepath.IsAbs(dest) {
			url = storage.NewFileURI(dest)
		} else {
			url = utils.GetImageURI(uri, dest)
		}
		return []widget.RichTextSegment{&widget.ImageSegment{Source: url, Title: string(t.Title), Alignment: fyne.TextAlignCenter}}, nil
	}
	return nil, nil
}

func suffixSpaceIfAppropriate(text string, n ast.Node) string {
	next := n.NextSibling()
	if next != nil && next.Type() == ast.TypeInline && !strings.HasSuffix(text, " ") {
		return text + " "
	}
	return text
}

func renderChildren(source []byte, n ast.Node, blockquote bool) ([]widget.RichTextSegment, error) {
	children := make([]widget.RichTextSegment, 0, n.ChildCount())
	for childCount, child := n.ChildCount(), n.FirstChild(); childCount > 0; childCount-- {
		segs, err := renderNode(source, child, blockquote)
		if err != nil {
			return children, err
		}
		children = append(children, segs...)
		child = child.NextSibling()
	}
	return children, nil
}

func forceIntoText(source []byte, n ast.Node) string {
	texts := make([]string, 0)
	ast.Walk(n, func(n2 ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch t := n2.(type) {
			case *ast.Text:
				texts = append(texts, string(t.Text(source)))
			}
		}
		return ast.WalkContinue, nil
	})
	return strings.Join(texts, " ")
}

func forceIntoHeadingText(source []byte, n ast.Node) string {
	var text strings.Builder
	ast.Walk(n, func(n2 ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			switch t := n2.(type) {
			case *ast.Text:
				text.Write(t.Text(source))
			}
		}
		return ast.WalkContinue, nil
	})
	return text.String()
}

// Parse :: parse markdown and returns richText segments
func parse(content []byte) []widget.RichTextSegment {
	r := mdRenderer{}
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
			extension.Strikethrough,
			extension.Linkify,
		),
		goldmark.WithRenderer(&r),
	)
	err := md.Convert(content, nil)
	if err != nil {
		fyne.LogError("Failed to parse markdown", err)
	}
	return r
}

type Md struct {
	*widget.RichText
}

func (m *Md) Hello() {
	fmt.Println("Hello")
	// m.RichText
}

var uri fyne.URI

func New(URI fyne.URI, content []byte) *Md {
	uri = URI
	return &Md{
		widget.NewRichText(parse(content)...),
	}
}
