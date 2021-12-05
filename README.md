This is an example project I wrote for a youtube tutorial about webscraping using golang and [gocolly](https://github.com/gocolly/colly)

It extracts data from a [tracking differences](https://www.trackingdifferences.com/) website and stores it into a struct.

```go
type EtfInfo struct {
	Title              string
	Replication        string
	Earnings           string
	TotalExpenseRatio  string
	TrackingDifference string
	FundSize           string
}
```