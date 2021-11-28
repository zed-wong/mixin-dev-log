package main

import(
	"fmt"
	"time"
	"github.com/tidwall/gjson"
)
var (
	TimeLayout = "2006-01-02T15:04:05Z"
	PandoLeafVatsEndpoint="https://leaf-api.pando.im/api/vats/"
        PandoLeafCatsEndpoint="https://leaf-api.pando.im/api/cats"
        PandoLeafOraclesEndpoint="https://leaf-api.pando.im/api/oracles"
        PandoLeafHistoryEndpoint="https://leaf-api.pando.im/api/vats/%s/events"
	myvaultid = "ca7d8a0f-2c4e-3027-8751-34e2d6abe335"
)

func calcRatio(vault_id string){
	curl1 := HttpGet(PandoLeafVatsEndpoint+vault_id)
	ink := gjson.Get(curl1, `data.ink`).Float()
	art := gjson.Get(curl1,`data.art`).Float()
	collateral_id := gjson.Get(curl1,`data.collateral_id`).String()

        curl2 := HttpGet(PandoLeafCatsEndpoint)
        catpatt := fmt.Sprintf(`data.collaterals.#(id="%s")`,collateral_id)
        collateral := gjson.Get(curl2,catpatt)
        gem := collateral.Get("gem").String()
        mat := collateral.Get("mat").Float()
        duty := collateral.Get("duty").Float()

        curl3 := HttpGet(PandoLeafOraclesEndpoint)
        orapatt := fmt.Sprintf(`data.oracles.#(asset_id="%s")`, gem)
        oracle := gjson.Get(curl3,orapatt)
        current := oracle.Get("current").Float()
        next := oracle.Get("next").Float()

        vaultHistoryURL := fmt.Sprintf(PandoLeafHistoryEndpoint,vault_id)
        data := HttpGet(vaultHistoryURL)
	eventArray := gjson.Get(data, `data.events`)

        event := gjson.Get(data, `data.events.#(dart!=0)`)
        event_created_at := event.Get(`created_at`).String()
        event_dart := event.Get(`dart`).Float()
        event_debt := event.Get(`debt`).Float()

	if gjson.Get(data, `data.events.0.dart`).Float() == 0{
		eventArray.ForEach(func(key,value gjson.Result) bool{
			dink := value.Get("dink").Float()
			ink += dink
			if value.Get(`created_at`).String() == event_created_at{
				return false
			}
			return true
		})
	}

        var ratio, next_ratio, close_out_price float64

        daysTilToday := func(created_at string)float64{
                now := time.Now()
                then, err := time.Parse(TimeLayout, created_at)
                if err !=nil {
                        fmt.Println(err)
                }
                diff := now.Sub(then)
                return diff.Hours()/24
        }
        calcCompound := func(in,days,rate float64)float64{
                for i:=0; i<int(days); i++{
                        in = in * rate/365 + in
                }
                return in
        }
        event_debt_now := event_debt/event_dart*art
        interest_rate := duty-1

        event_days := daysTilToday(event_created_at)
        debt := calcCompound(event_debt_now, event_days, interest_rate)

        if debt != 0{
                ratio            = ink*current/debt
                next_ratio       = ink*next/debt
                close_out_price  = mat/ink*debt
        } else {
                ratio            = 0
                next_ratio       = 0
                close_out_price  = 0
        }

	fmt.Println(ratio,"\n", next_ratio,"\n", close_out_price,"\n")
//	fmt.Printf("%T\n", eventArray)
}

func main(){
	calcRatio(myvaultid)
}
