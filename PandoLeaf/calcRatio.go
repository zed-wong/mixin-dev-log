package main

import(
	"fmt"
	"log"
	"github.com/tidwall/gjson"
	"github.com/go-resty/resty/v2"
)
type VAULT struct{
        UserID  string `json:"user_id"`
        VaultID string `json:"vault_id"`
        IdentityID string `json:"identity_id"`
        Avatar  string `json:"avatar"`
        Ratio   string `json:"ratio"`
        NextRatio string `json:"next_ratio"`
        AlertRatio      string `json:"alert_ratio"`
        AddAt   string `json:"add_at"`
        EndAt   string `json:"end_at"`
        Triggered       bool `json:"triggered"`
        Free    bool `json:"free"`
}

func FetchVaultData(vaultID string) *VAULT{
        client := resty.New()
        respvat, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/vats/%s", vaultID))
        if err != nil {
                log.Println(err)
        }
        vault := gjson.Get(respvat.String(), "data")
        art := vault.Get("art").Float()         // art*rate = debt
        ink := vault.Get("ink").Float()         // Amount of coll(BTC)
        catID := vault.Get("collateral_id").String()
        identityID := vault.Get("identity_id").String()

        respcat, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/cats/%s", catID))
        if err != nil {
                log.Println(err)
        }
        collateral := gjson.Get(respcat.String(), "data")
        rate := collateral.Get("rate").Float()  // art*rate = debt
        price := collateral.Get("price").Float()// Current price
        mat := collateral.Get("mat").Float()    // Minium ratio
        gem := collateral.Get("gem").String()   // coll (BTC) UUID

        respasset, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/assets/%s", gem))
        if err != nil {
                log.Println(err)
        }
        avatar := gjson.Get(respasset.String(), "data.logo").String()   // Avatar url

        oracleresp, err := client.R().Get(fmt.Sprintf("https://leaf-api.pando.im/api/oracles/%s", gem))
        if err != nil {
                log.Println(err)
        }
        nextPrice := gjson.Get(oracleresp.String(), "data.next").Float()

        ratio := fmt.Sprintf("%.2f", (ink*price)/(art*rate)*100)
        nextRatio := fmt.Sprintf("%.2f", (ink*nextPrice)/(art*rate)*100)
        alertRatio := fmt.Sprintf("%.0f", mat*1.5*100)

        return &VAULT{
                VaultID: vaultID,
                IdentityID: identityID,
                Avatar: avatar,
                Ratio: ratio,
                NextRatio: nextRatio,
                AlertRatio: alertRatio,
                Triggered: false,
        }
}

