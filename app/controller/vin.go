package controller

import (
	"github.com/gin-gonic/gin"
)

type VinByPlateRequest struct {
	Plate string `json:"plate" example:"M343TT123" validate:"required"`
}

// VinByPlate godoc
// @Summary      Распознать vin и другие данные по госномеру
// @Description  Ести метод вызвать 2 раза, то он потратит 1 сканирование. Вернет из кэша. 
// @Tags         Распознание
// @Accept       json
// @Produce      json
// @Param 		 vin body VinByPlateRequest true "Распознать по госномеру"
// @Success      200  {object}   model.Vin
// @Failure      400  {object}  controller.ActionError
// @Router       /vin [post]
// @Security 	 ApiKeyAuth
func (c *Config) VinByPlate(g *gin.Context) {
	req := &VinByPlateRequest{}
	user, err := c.initRequest(g, req)
	if err != nil {
		c.error(g, err)
		return
	}

	new, err := c.service.Vin().VinByPlate(req.Plate, user.ID, true)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, new)
}

type VinByPlateBulkRequest struct {
	Plates []string `json:"plate" example:"M343TT123,B345KY24" validate:"required"`
}

// VinByPlateBulk godoc
// @Summary      Распознать vin и другие данные по госномеру пачкой
// @Description  Ести метод вызвать 2 раза, то он потратит 1 сканирование. Вернет из кэша. Максимум 10000 номеров за раз.
// @Description  Распознание происходит не сразу, отложенно (12 номеров в минуту).
// @Tags         Распознание
// @Accept       json
// @Produce      json
// @Param 		 vin body VinByPlateBulkRequest true "Распознать по госномерам"
// @Success      200  {array}   model.Vin
// @Failure      400  {object}  controller.ActionError
// @Router       /vin/bulk [post]
// @Security 	 ApiKeyAuth
func (c *Config) VinByPlateBulk(g *gin.Context) {
	req := &VinByPlateBulkRequest{}
	user, err := c.initRequest(g, req)
	if err != nil {
		c.error(g, err)
		return
	}

	news, err := c.service.Vin().VinByPlateBulk(req.Plates, user.ID)
	if err != nil {
		c.error(g, err)
		return
	}

	c.respond(g, news)
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
