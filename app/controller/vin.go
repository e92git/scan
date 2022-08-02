package controller

import (
	"github.com/gin-gonic/gin"
)

type VinByPlateRequest struct {
	Place     string `json:"place" example:"pokrovka"`
	Plate     string `json:"plate" example:"M343TT123"`
	ScannedAt string `json:"scanned_at" example:"2022-07-23 11:23:55"`
}

// VinByPlate godoc
// @Summary      Распознать vin и другие данные по госномеру
// @Tags         Распознание
// @Accept       json
// @Produce      json
// @Param 		 scan body VinByPlateRequest true "Распознать по госномеру"
// @Success      200  {array}   model.Vin
// @Failure      400  {object}  controller.ActionError
// @Router       /scan [post]
// @Security 	 ApiKeyAuth
func (c *Config) VinByPlate(g *gin.Context) {
	scan := &AddScanRequest{}
	if err := g.BindJSON(scan); err != nil {
		c.error(g, err)
		return
	}

	newScan, err := c.service.Scan().Create(
		scan.Place,
		scan.Plate,
		scan.ScannedAt,
	)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, newScan)
}

// Уже был найден ранее
// {
//     "state": "ok",
//     "size": 1,
//     "version": "2.0",
//     "stamp": "2022-08-02T14:36:18.810Z",
//     "data": [
//         {
//             "uid": "report_check_vehicle_eyJ0eXBlIjoiR1JaIiwiYm9keSI6ItCeMjQ10JrQnDQyIiwic2NoZW1hX3ZlcnNpb24iOiIxLjAiLCJzdG9yYWdlcyI6e319_i_ec8a0f22e8b235729192493e785b9244@e92",
//             "isnew": false,
//             "suggest_get": "2022-08-02T11:53:54.782Z"
//         }
//     ]
// }

/// В процессе поиска
// {
//     "state": "ok",
//     "size": 1,
//     "version": "2.0",
//     "stamp": "2022-08-02T11:53:53.310Z",
//     "data": [
//         {
//             "uid": "report_check_vehicle_eyJ0eXBlIjoiR1JaIiwiYm9keSI6ItCeMjQ10JrQnDQyIiwic2NoZW1hX3ZlcnNpb24iOiIxLjAiLCJzdG9yYWdlcyI6e319_i_ec8a0f22e8b235729192493e785b9244@e92",
//             "isnew": true,
//             "make_mode_log": {
//                 "make_mode": "TRANSACTIONAL_CONDITIONAL_NON_LOCK",
//                 "need_transaction": true,
//                 "balance_map": {
//                     "TOTAL": {
//                         "quote_init": 0,
//                         "quote_up": 0,
//                         "quote_use": 11
//                     },
//                     "MONTH": {
//                         "quote_init": 0,
//                         "quote_up": 0,
//                         "quote_use": 0
//                     },
//                     "DAY": {
//                         "quote_init": 0,
//                         "quote_up": 0,
//                         "quote_use": 0
//                     }
//                 },
//                 "need_lock": false,
//                 "need_balance_calculate": true,
//                 "need_priority_calculate": false
//             },
//             "process_request_uid": "report_check_vehicle_eyJ0eXBlIjoiR1JaIiwiYm9keSI6ItCeMjQ10JrQnDQyIiwic2NoZW1hX3ZlcnNpb24iOiIxLjAiLCJzdG9yYWdlcyI6e319_i_ec8a0f22e8b235729192493e785b9244_ab651ec3a38f41b68065bd3de9e81216@e92",
//             "suggest_get": "2022-08-02T11:53:53.260Z"
//         }
//     ]
// }
