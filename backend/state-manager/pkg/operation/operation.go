package operation

import (
	"fmt"
	statev1 "github.com/ueckoken/plarail2023/backend/spec/state/v1"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/db"
	"github.com/ueckoken/plarail2023/backend/state-manager/pkg/mqtt_handler"

	"log"
)

// シンプルなオペレーション用

func Check(change *statev1.StopAndState) {
	defer db.C()
	db.Open()
	state, _ := db.GetBlock("yamashita_b1")
	fmt.Println(state.State.String())
	if state.State == statev1.BlockStateEnum_BLOCK_STATE_CLOSE {
		// 2. 閉塞が閉じていたらストップレールをあげる
		err := db.UpdateStop(&statev1.StopAndState{
			Id:    "yamashita_s1",
			State: statev1.StopStateEnum_STOP_STATE_STOP,
		})
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		// 3. 閉塞が開いていたらストップレールを下げる
		err := db.UpdateStop(&statev1.StopAndState{
			Id:    "yamashita_s1",
			State: statev1.StopStateEnum_STOP_STATE_GO,
		})
		if err != nil {
			log.Fatalln(err)
		}
		mqtt_handler.NotifyStateUpdate("stop", "yamashita_s1", statev1.StopStateEnum_STOP_STATE_STOP.String())
	}
}
