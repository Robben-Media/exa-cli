package exa

// SearchRequest represents a search request to the Exa API.
type SearchRequest struct {
	Query          string          `json:"query"`
	NumResults     int             `json:"numResults,omitempty"`
	Type           string          `json:"type,omitempty"`
	UseAutoprompt  bool            `json:"useAutoprompt,omitempty"`
	IncludeDomains []string        `json:"includeDomains,omitempty"`
	ExcludeDomains []string        `json:"excludeDomains,omitempty"`
	StartPublished string          `json:"startPublishedDate,omitempty"`
	EndPublished   string          `json:"endPublishedDate,omitempty"`
	Contents       *ContentOptions `json:"contents,omitempty"`
}

// ContentOptions controls what content is returned with search results.
type ContentOptions struct {
	Text       *TextOptions `json:"text,omitempty"`
	Highlights bool         `json:"highlights,omitempty"`
	Summary    bool         `json:"summary,omitempty"`
}

// TextOptions controls text content retrieval.
type TextOptions struct {
	MaxCharacters int  `json:"maxCharacters,omitempty"`
	IncludeHTML   bool `json:"includeHtmlTags,omitempty"`
}

// SearchResponse is the response from a search request.
type SearchResponse struct {
	Results        []SearchResult `json:"results"`
	AutopromptText string         `json:"autopromptString,omitempty"`
}

// SearchResult represents a single search result.
type SearchResult struct {
	URL           string   `json:"url"`
	Title         string   `json:"title"`
	Score         float64  `json:"score,omitempty"`
	PublishedDate string   `json:"publishedDate,omitempty"`
	Author        string   `json:"author,omitempty"`
	Text          string   `json:"text,omitempty"`
	Highlights    []string `json:"highlights,omitempty"`
	Summary       string   `json:"summary,omitempty"`
}

// FindSimilarRequest represents a find-similar request.
type FindSimilarRequest struct {
	URL            string          `json:"url"`
	NumResults     int             `json:"numResults,omitempty"`
	IncludeDomains []string        `json:"includeDomains,omitempty"`
	ExcludeDomains []string        `json:"excludeDomains,omitempty"`
	Contents       *ContentOptions `json:"contents,omitempty"`
}

// AnswerRequest represents an answer request.
type AnswerRequest struct {
	Query      string `json:"query"`
	NumResults int    `json:"numResults,omitempty"`
	Text       bool   `json:"text,omitempty"`
}

// AnswerResponse is the response from an answer request.
type AnswerResponse struct {
	Answer    string         `json:"answer"`
	Citations []SearchResult `json:"citations,omitempty"`
}
