package cataloginventory

import (
	"time"

	"github.com/gocql/gocql"
)

type CatalogInventoryStockItem struct {
	ItemID                  gocql.UUID `json:"item_id"`
	ProductID               gocql.UUID `json:"product_id"`
	StockID                 gocql.UUID `json:"stock_id"`
	Qty                     float64    `json:"qty"`
	MinQty                  float64    `json:"min_qty"`
	UseConfigMinQty         int        `json:"use_config_min_qty"`
	IsQtyDecimal            int        `json:"is_qty_decimal"`
	Backorders              int        `json:"backorders"`
	UseConfigBackorders     int        `json:"use_config_backorders"`
	MinSaleQty              float64    `json:"min_sale_qty"`
	UseConfigMinSaleQty     int        `json:"use_config_min_sale_qty"`
	MaxSaleQty              float64    `json:"max_sale_qty"`
	UseConfigMaxSaleQty     int        `json:"use_config_max_sale_qty"`
	IsInStock               int        `json:"is_in_stock"`
	LowStockDate            *time.Time `json:"low_stock_date"`
	NotifyStockQty          float64    `json:"notify_stock_qty"`
	UseConfigNotifyStockQty int        `json:"use_config_notify_stock_qty"`
	ManageStock             int        `json:"manage_stock"`
	UseConfigManageStock    int        `json:"use_config_manage_stock"`
	StockStatusChangedAuto  int        `json:"stock_status_changed_auto"`
	UseConfigQtyIncrements  int        `json:"use_config_qty_increments"`
	QtyIncrements           float64    `json:"qty_increments"`
	UseConfigEnableQtyInc   int        `json:"use_config_enable_qty_inc"`
	EnableQtyIncrements     int        `json:"enable_qty_increments"`
	IsDecimalDivided        int        `json:"is_decimal_divided"`
	WebsiteID               gocql.UUID `json:"website_id"`
}
