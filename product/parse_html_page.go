package product

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

var reWarehouseInventory = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblWhsInv"[^>]*>([^<]+?)</span>`)
var reWarehouseOnOrder = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblWhsOnOrder"[^>]*>([^<]+?)</span>`)
var reSku = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblSku"[^>]*>([^<]+?)</span>`)
var reName = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblDesc"[^>]*><b>([^<]+?)</b></span>`)
var reStatusDescription = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblStatus"[^>]*>([^<]+?)</span>`)
var reCurrentRetail = regexp.MustCompile(`<span [^>]*id="ContentPlaceHolderBody_lblPrice"[^>]*>([^<]+?)</span>`)
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

func (p *Product) ParseProductPage(body []byte) error {
	var err error

	if inventory := reWarehouseInventory.FindSubmatch(body); len(inventory) > 1 {
		p.WarehouseInventory, err = strconv.Atoi(string(inventory[1]))
		if err != nil {
			return err
		}
	}

	if onOrder := reWarehouseOnOrder.FindSubmatch(body); len(onOrder) > 1 {
		p.WarehouseOnOrder, err = strconv.Atoi(string(onOrder[1]))
		if err != nil {
			return err
		}
	}

	if sku := reSku.FindSubmatch(body); len(sku) > 1 {
		p.CSCode = string(sku[1])
	}

	if name := reName.FindSubmatch(body); len(name) > 1 {
		p.Name = string(name[1])
	}

	if statusDescription := reStatusDescription.FindSubmatch(body); len(statusDescription) > 1 {
		p.StatusDescription = string(statusDescription[1])
		if p.StatusCode == "" {
			p.StatusCode = p.StatusDescriptionToCode()
		}
	}

	if currentRetail := reCurrentRetail.FindSubmatch(body); len(currentRetail) > 1 {
		p.CurrentRetail, err = strconv.ParseFloat(strings.Replace(string(currentRetail[1]), "$", "", 1), 64)
		if err != nil {
			return err
		}
	}

	storesInventory := 0

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
				return err
			}
			storesInventory += s.Qty

			stores = append(stores, s)
		}

		p.Stores = stores
	}

	p.StoresInventory = storesInventory
	p.InventoryUpdated = time.Now()

	return nil
}
