package astichartjs

// Chart types
const (
	ChartTypeBar     = "bar"
	ChartTypeLine    = "line"
	ChartTypePie     = "pie"
	ChartTypeScatter = "scatter"
)

// Chart axis positions
const (
	ChartAxisPositionsBottom = "bottom"
)

// Chart axis type
const (
	ChartAxisTypesLinear      = "linear"
	ChartAxisTypesTime        = "time"
	ChartAxisTypesTimeseries  = "timeseries"
	ChartAxisTypesLogarithmic = "logarithmic"
	ChartAxisTypesCategory    = "category"
)

// Chart background colors
const (
	ChartBackgroundColorBlue   = "rgba(54, 162, 235, 0.2)"
	ChartBackgroundColorGreen  = "rgba(75, 192, 192, 0.2)"
	ChartBackgroundColorOrange = "rgba(255, 159, 64, 0.2)"
	ChartBackgroundColorPurple = "rgba(153, 102, 255, 0.2)"
	ChartBackgroundColorRed    = "rgba(255, 99, 132, 0.2)"
	ChartBackgroundColorYellow = "rgba(255, 206, 86, 0.2)"
)

// Chart border colors
const (
	ChartBorderColorBlue   = "rgba(54, 162, 235, 1)"
	ChartBorderColorGreen  = "rgba(75, 192, 192, 1)"
	ChartBorderColorOrange = "rgba(255, 159, 64, 1)"
	ChartBorderColorPurple = "rgba(153, 102, 255, 1)"
	ChartBorderColorRed    = "rgba(255,99,132,1)"
	ChartBorderColorYellow = "rgba(255, 206, 86, 1)"
)

// Chart represents a chart
type Chart struct {
	Data    *Data    `json:"data,omitempty"`
	Options *Options `json:"options,omitempty"`
	Type    string   `json:"type"`
}

// Data represents data
type Data struct {
	Datasets []Dataset `json:"datasets,omitempty"`
	Labels   []string  `json:"labels,omitempty"`
}

// Dataset represents a dataset
type Dataset struct {
	BackgroundColor interface{}   `json:"backgroundColor,omitempty"`
	BorderColor     interface{}   `json:"borderColor,omitempty"`
	Data            []interface{} `json:"data,omitempty"`
	Label           string        `json:"label,omitempty"`
	PointRadius     string        `json:"pointRadius,omitempty"`
	Stack           string        `json:"stack,omitempty"`
}

// DataPoint represents a data point
type DataPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// TimedDataPoint represents a time-based data point
type TimedDataPoint struct {
	T string  `json:"x"`
	Y float64 `json:"y"`
}

// Options represents chart options
type Options struct {
	Animation           *Animation `json:"animation,omitempty"`
	Elements            *Elements  `json:"elements,omitempty"`
	Hover               *Hover     `json:"hover,omitempty"`
	MaintainAspectRatio *bool      `json:"maintainAspectRatio,omitempty"`
	Responsive          *bool      `json:"responsive,omitempty"`
	Scales              *Scales    `json:"scales,omitempty"`
	Plugins             *Plugins   `json:"plugins,omitempty"`
}

// Plugins represents plugin options
type Plugins struct {
	Legend *Legend `json:"legend,omitempty"`
	Title  *Title  `json:"title,omitempty"`
}

// Scales represents scales options (Chart.js 4.x)
type Scales struct {
	X Axis `json:"x,omitempty"`
	Y Axis `json:"y,omitempty"`
}

// Axis represents axis options (Chart.js 4.x)
type Axis struct {
	Display  *bool      `json:"display,omitempty"`
	Position string     `json:"position,omitempty"`
	Stacked  *bool      `json:"stacked,omitempty"`
	Ticks    *Ticks     `json:"ticks,omitempty"`
	Type     string     `json:"type,omitempty"`
	TimeType *TimeType  `json:"time,omitempty"`
	Title    *AxisTitle `json:"title,omitempty"`
}

// AxisTitle represents a axis title options
type AxisTitle struct {
	Display *bool       `json:"display,omitempty"`
	Text    interface{} `json:"text,omitempty"` // string or []string
	Color   string      `json:"color,omitempty"`
	Font    *Font       `json:"font,omitempty"`
	Padding *Padding    `json:"padding,omitempty"`
	Align   string      `json:"align,omitempty"` // "start", "center", "end"
}

// Padding ...
type Padding struct {
	Top    *int `json:"top,omitempty"`
	Bottom *int `json:"bottom,omitempty"`
	Left   *int `json:"left,omitempty"`
	Right  *int `json:"right,omitempty"`
}

// TimeType ...
type TimeType struct {
	DisplayFormats *TimeDisplayFormats `json:"displayFormats,omitempty"`
}

// TimeDisplayFormats ...
type TimeDisplayFormats struct {
	TimeDisplayFormatMillisecond string `json:"millisecond,omitempty"`
}

// Ticks represents tick options (Chart.js 4.x)
type Ticks struct {
	Font *Font `json:"font,omitempty"`
}

// Font represents tick font options
type Font struct {
	Size       *int    `json:"size,omitempty"`
	Family     string  `json:"family,omitempty"`
	Weight     string  `json:"weight,omitempty"`
	Style      string  `json:"style,omitempty"`
	LineHeight float64 `json:"lineHeight,omitempty"`
}

// Title represents a chart title
type Title struct {
	Display *bool  `json:"display,omitempty"`
	Text    string `json:"text,omitempty"`
	Font    *Font  `json:"font,omitempty"`
}

// Legend represents legend options
type Legend struct {
	Display *bool         `json:"display,omitempty"`
	Labels  *LegendLabels `json:"labels,omitempty"`
}

// LegendLabels represents legend label options
type LegendLabels struct {
	Font          *Font `json:"font,omitempty"`
	UsePointStyle *bool `json:"usePointStyle,omitempty"`
}

// Animation represents animation options
type Animation struct {
	Duration *int `json:"duration,omitempty"`
}

// Hover represents hover options (Chart.js 4.x — no animationDuration)
type Hover struct{}

// Elements represents element options
type Elements struct {
	Line *Line `json:"line,omitempty"`
}

// Line represents line element options
type Line struct {
	Fill    *bool    `json:"fill,omitempty"`
	Tension *float64 `json:"tension,omitempty"`
}
