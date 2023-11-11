package nihonshu

type (
	NIHONSHU_ID uint64
	BREWERY_ID  uint64
)

type Nihonshu struct {
	ID        NIHONSHU_ID `json:"id"`
	Name      string      `json:"name"`
	BreweryID BREWERY_ID  `json:"BreweryID"`
}
