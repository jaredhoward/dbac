package product

import (
	"regexp"
	"strconv"
)

var reWarehouseInventory = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblWhsInv"[^>]*>([^<]+?)</span>`)
var reWarehouseOnOrder = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblWhsOnOrder"[^>]*>([^<]+?)</span>`)
var reInventoryDetailsTable = regexp.MustCompile(`(?ims)<table [^>]*id="ContentPlaceHolderBody_gvInventoryDetails"[^>]*>(.+?)</table>`)
var reInventoryDetailsRow = regexp.MustCompile(`(?ims)` +
	`<tr [^>]*?class="gridViewRow"[^>]*?>.*?` +
	`<td>(?P<id>.*?)</td>` +
	`.*?` +
	`<td>(?P<name>.+?)</td>` +
	`.*?` +
	`<td[^>]*?>(?:.*?<span[^>]*?>)?(?P<qty>.+?)(?:</span>.*?)?</td>` +
	`.*?` +
	`<td>(?P<address>.+?)</td>` +
	`.*?` +
	`<td>(?P<city>.+?)</td>` +
	`.*?` +
	`<td>(?P<phone>.+?)</td>` +
	`.*?</tr>`)

func ParseProductPage(body []byte) (*Product, error) {
	var err error
	p := &Product{}

	inventory := reWarehouseInventory.FindSubmatch(body)
	p.WarehouseInventory, err = strconv.Atoi(string(inventory[1]))
	if err != nil {
		return nil, err
	}

	onOrder := reWarehouseOnOrder.FindSubmatch(body)
	p.WarehouseOnOrder, err = strconv.Atoi(string(onOrder[1]))
	if err != nil {
		return nil, err
	}

	detailsTable := reInventoryDetailsTable.FindSubmatch(body)
	if detailsTable != nil {
		stores := make([]store, 0)

		for _, row := range reInventoryDetailsRow.FindAllSubmatch(detailsTable[1], -1) {
			result := make(map[string]string)
			for i, name := range reInventoryDetailsRow.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = string(row[i])
				}
			}

			s := store{
				ID:      result["id"],
				Name:    result["name"],
				Address: result["address"],
				City:    result["city"],
				Phone:   result["phone"],
			}
			s.Qty, err = strconv.Atoi(result["qty"])
			if err != nil {
				return nil, err
			}
			stores = append(stores, s)
		}

		p.Stores = stores
	}

	return p, nil
}
