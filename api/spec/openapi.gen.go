// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/trustbloc/vcs/pkg/restapi/v1/common"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+y963IbN/Yg/ioo/v9VsWtJys5l5hfvl58iKQkzTqSRZLmmYhUL6gZJWM1GB0CL5qS0",
	"ta+xr7dPsoUDoBvoRt8oUXEm+pRYbNwOzjk49/P7KGLrjKUklWL05veRiFZkjeF/D6OICHHJbkl6TkTG",
	"UkHUn2MiIk4zSVk6ejP6mcUkQQvGkf4cwffIDpiOxqOMs4xwSQnMiuGzuVSf1ae7XBGkv0DwBaJC5CRG",
	"N1sk1U+5XDFO/43V50gQfke4WkJuMzJ6MxKS03Q5uh+PvA/nMZGYJqK+3PnJP9/Nzk+O0WZFUhQchDLM",
	"8ZpIwhEVKBckRpIhTn7LiZCwPZxGBLEFwigiXGKaoiNOYpJKihOkdoawQDFZ0JTEiKbogkSw/W+mr6ev",
	"p2gm0c/vLi7RL6eX6IboFZhcEb6hgsDPVCCcIsw53qp12M1HEkkxbpj27+qbX8+/P/r2q2//dq2gQyVZ",
	"w+H/f04Wozej6UHE1muWTrd4nfx/ByUCHJjbPzh0IXFsoHdfwBm2ov4dzVOWRgG0uICbQBFLFUDU/2IE",
	"nyrg2VNKhiJOsCQIo4wzdbQFypgQRAh1ErZAt2SL1lgSrmAJl2Qgr6eMCkAHscBsb04+ZZQTMacBjJul",
	"kiwJRzFJGcyq8CyhCyLpmii4ChKxNBZqN+onM6ezHtUzqAXbFrpsn9fF+vDknCw4Eas20jGf6FnGaLOi",
	"0QpFOHVBzm4AR1Oy8dYUQQiKiGWB6z09u5yd/nL4dozoAlG4gkghO4OjwCB7USXxRgklqfyfJXKPkaW/",
	"4Nqwrbn+c+iwQFoGei6zCEwG0Pstp5zEoze/+jzIW+h6PJJUJmpsiP0VE2saHI1HnyYSL4WalNE4+jqi",
	"o+v78egwuj3hnPFmvnkY3SLeyCSJGlwfBHMi52/dR9Uzece63eU45/o2mw4CP9bPUdKnYb4ZJxGWamuS",
	"52Rcmeu4+H2KnHkdMhcoxhKHmFqYjxUbn0myDnEwckfSwKkuHdxSfGFBI/3mwPdBdIVf5t401Vl/zNc4",
	"nXCCY3yTEHR4cTSbIUk+ScX+7mgMTC2OqfocJ4imC8bXsO64IF8sBBUSNuY8MzOF+Qo17kiiLkAxmDyN",
	"CRcSp7Fla7BFJFdYIhZFOedBYhkD0+E4qr6dPrKo72DVuWYAC0oCKHua2cPoHZbfBld2YT2ncRjdZsdh",
	"vPdQXKPrcAwHRBmOq+7oGjWHccydYT/45q7wp8G9bpxyj/XY+FVlQU245vLY6qSWq1yP6yTzHZbRqgRe",
	"I1st5erT2fERulHDXKD3YLlz8w38uRe7rO+rxjErRw+t5jw2DaftS5W14d1aCEDruzq0ml7aRgn2p4vT",
	"X5B4GjH26OFiLGyXPqYs612tBp+PSSwlp4vRm19/r+24P5bpeSv3PLq/HoR3dnNtiDdQ4imHHrF0QZc5",
	"B+oWF3mWMQ6PQu1lSo1mppmc/vGGCCQyEin+UIDdVQ/Vp2F+KvRSwtUxA/ibYLoOaLbfM47Wgs3XMYsQ",
	"TmN0F/0PEU8+biS6ixBLk+0UnertetidKAbPFijFa3Jwh5OcoAxTLpQyQThBBEcr+LHktEIpYmobCN+w",
	"XB9H5HputlgQrvVT/5RTpER4vYBRUHAKmgESebSyoHyRahVCSX2KGvNI5pyIl2PEuKcUO4NcTaZkvA7G",
	"gNJM7ZPZWykuN39cTuDPLOhSwXGOk+UczibmogVj7OYjLAgSJBVU0jtiuI7QyGHAbOwfyZJxKldrUWKO",
	"QZdcEKXJIbUF+LuxnPi8pSDeurZVVe35NpNsyXG2otH8hsJLPl8TuWLxI55qxTZV/KcC3bA8ja06WT7p",
	"loBO0njyThCONitmOa06vY9hg44bU5EleBsk67rlxaEF5hGR3oSZDJWkandewM3RaeDdKo1HCU6XOV6S",
	"3kqOg5fmEKHzsSisSXuMomANxn5jr8m+JRXDVtUE9evs4nT6+r9evf5q8s118CnTQmUAysh9b6vL6lEa",
	"hlQ4oBsjOiXTMfq4kfO7aP5RqOeWoyTO5nfRFB2TjGgJlKXuRECaY/hL9foWOQcmRBKyVlDWx7Mb0da8",
	"NEYvmJE7k+1LlGEuaZQnmGs+qJHAueCfD/9lV4DRjnBteCaQASsQxx8fhCTjcUg2LqhPW1wUVwZurbmR",
	"Jj7F42GPa8uXYTL1f1skVixPYsWPzWZKA857nCREDqMrEIjAtlJhGqWuceY9aG2YfqYmU6pS+QzfV/Wy",
	"035vsJLIYG8vxMs+r3DwTWmwjrUjs7aO6ZfPLExF2/uv2AN84+JZO3LcRTJM6QEpwJB6TNTLgaWH6mDV",
	"PnLIzaf3lZSZeHNwoF5nyXF0S/iUErmYMr48iFl0sJLr5CDmeCEn6u8ThnO5mugdTO6iyavXncqV4RiO",
	"bNcpm1miLt/5aavgp1XHitx3XD4IvsR1g6PbJVcP1DxiiTbT1S4gYRFOSMNPS9aF6G/VN0pdxevwJEpx",
	"b1k+50ng7/chGNpzNgCoET4zI5X+SIVkfHuMJa6jXOvniCtdWwCXrTDMQuRd6c/NE2yYcqvSG1LqXeIK",
	"25qdCYBXNShYhSQQ+Q+hGMYUQZEzXiYsAxzkpPgAHWNJGg0lCkYNU1iAt08QekJmvSwpGWcLmpD5HeEi",
	"aHwy05zp75D5Lmzp5zgVxt4Yur/L8vdeBhkfHYqTBq45yFYquFpYD4YzkXPtjzm8wzTBNwnpY8FwkPVd",
	"pu62xZl6RzhdUDXzmaYkwBnHqNTGZK5aB1dh2r5UEI56+426dwVS/QxhA01g+1P12gyU5kl1NR2tJlc1",
	"f6XqEPVp6SzThiXjyULvVyQtHn/fDT12JdryVyVf4nSrvWzuguZLKwmVQ4TnfzYsuYtLWpqYkxQ0RR/C",
	"Pe0+J+XYFt3ge0f6914JDbpGp58RPru29dP7S5ArG97HoTbLHcyVvQyVOIpIJoHhOw7gAT4K1GjR1M5T",
	"kd8IddZUJtuqs9gzUWp0KVFFGzS91xulTCJOZM7Thqvx7K79T/EnsMoOPMyebLZD9+LgIKgmav42A2+F",
	"lV97iwfEpyOrYHuRLKlSPdGa8YC5SuwWulKnv/Ivp8XDU5XIXN6zGzZ697fwWFbQEjfcNRWwybgAdCBn",
	"V6QpfOHcrStK9NK7PJzy77i+5eswA93Z9t7kH2gT8I2OUPcBNUtRUdvMFm+DlziQQna5YpzWly4NFtok",
	"g2gaJXlMhLXn4Og2ZZuExEsNDEdUGix821ifoFEQo2OyIJyTGBUKhzPhFF2CwQ7sUOp/NDRLh4B97xBd",
	"NBhgNligPAWftWSIrtckpliSZKvB0uJWoKL1TbPLkwjM087KGypX8HNxNufHkzTOGA0KSI1aSDthVLF7",
	"dzo58WSxoF3MeTRdK6SSRawkV7f1tkRzJsuAqvz+BOFkWXorBkxfD3NIo/AKJI0eZ4WPm9s+4MJI0HSZ",
	"EJTlNwmNQHrASqj/6f0/NG7tvIcK4qgNjQG0+vit2OPc+WMgTouDsx2DtB17syKgd3S4NEulIeATVRpM",
	"I/cGSz7L1LDLtxchfOzteAv6PdVeFHb9ev790d+/ef23a3evjvvthUJwvdJL+/F/XTv+HWMz7zqXZSeK",
	"MZE0YnGVo6n3vRkaIJv/9P7SbuHb64GWqDR6Ingpcv2PgJc53Lyk2Cq4vmMsITg1z5BWuOG1bKcOM6E2",
	"hjqhWC6xuMhvHANhJoNm+m6KpxBk2a6VnaWAmd0Rvg3CUd2NOgpZKNHTkURAN9QRZcSd7pZsRT0KABnt",
	"ur7dBU6E2a+d+fBfKFoxQQowUhu75u8clmJc6aAOr73Rl1KPBQ5xjAbCCN9/T/b8KF6JC4llLloFYAGf",
	"1J9qUQxtwPLfO54lM4H5PHjqC++Tocc6zWRTAKn2gqmxoPl7Qrh/zH5n6TqC2krPU5QK4ZHSRgU5nR0f",
	"fX00O9pZF38nCDrQKxwYM7s4+N383+z4vvj/K21Tvz9wAnbFAWAXlmSi3vxJpDc1RaXRSf9JAdJstRWg",
	"bdrROd4gdeqESFKNaIBAFMUnolxItjbJJCErMI3nkqyzJOzHOA5Y/uznardpnoBt3cK17im/I5zTmMyb",
	"HB6n5gMTQ9oyacFEnFlNqNM8DipPdmpn8zY2KqZxv6UywpWcNVdHiqRiSzTGYSn/TH+K9Keo/LTPSo79",
	"swdSBy7y5FO0wumSeOlDRywmPaz7RI8F6SKXKwRP+4KztY3sBd9xIP6NklTOsRDqb6whL0Y/K/A22TgM",
	"uWFKEBBjJEiGOTYyCEYfRv/rwwhFK6wIinCtUS4oFxIEByqcZBaEpSRCu0LUr/rB0ha7li/P2Jn6OmxW",
	"rByoIQHmQpvxjbSgw7LKOO1crnROjiTeHrIssYHlJrgqlFGHXlwdXbzUB2dpsnWktOJ9/jDKefqGErl4",
	"A44E8Qbu541eaVJsf6K2/+bjRk7sLyUcPox0elsaw06dmDaz33UupH+YXLMthWDoy+krdFjONvkOq+Mf",
	"6aGH5Sh1MA2gNoAH/cZ6rtkxYOjV0YW21zvcNhyak83Vnno8Q8WXzlPUSUQ936WWeZr8EoV4t34oWTbm",
	"X+4vF1F+MnfY8fJ/qlgqu+HUE+A/EGn8vST2/EdtbG9JpNQOQDOy9S0unbDzzPHC1hcofbvIddeqGa2p",
	"e3SzlaTTFtG0ogPA5nO3Ac4cuBVyIns80OmLenc+C0rAzjF9eQenW23hvb8eACrPZl6u3ANoIusNtZmR",
	"86zhtSEu6iGR6j/niaRZUtMZsfFABWLR53EwEujcAApu7oyTiSU3xbIVT/k+YZtpyWMvCL+jEUE4kgJh",
	"gU7PYORG64LOQyaaBRsn+Bt2RoztIMToMV0j+7s9vdGOgdvpiF9HitM2bYhLX2FhPI+lGx8vpA5lj4gQ",
	"izxJtghHCgTASat5t50yrJHiu3zZPcS2aih8S/qYc+nuD+2BAdYdGvLkH6snvOItFk7EacRSQWPC1YXr",
	"eWKXYcVKqZF0TTq2YKPmGk8DH3REgRkNIxyPZH4MaSZO+AbarGhCfCSIGLhqtH2YCk+WKNKhx9YdYvQ8",
	"4zoBmtYSXq4eaUucAcVItGdI9mQdDzBZ9FzhqMTrJ+JRe9dmPy9aKJXfAB7bHwtLopJ0KUnAPVdOcqEV",
	"1im6sPZ7g2Y0XfbjXqH9PKYyHlpg/3q5s+ofoKI/HQ3bR0TTag9d3g404U16XIg+C/tvfwWiwtQNNRKB",
	"NopP3NI0hqh1/cIWPmSIMWZoSe/AjXx1dNGqC5r9z4sYWxNQ7S/+7vytG+IBBzJDIXXbESewTZ5Al/iW",
	"CKSeaQWNiCCFsEbhnW9IktymbFNEIpVRemAiv2FKBWvZpGZR1ckwh6xyay0H033q+N7tdRWnUCfb0CQp",
	"rCWa6zV8SdMi4CUjKY0nhQXSfvbm4KAN3sVO+xSS0SLgwYolwB0dkwZgmzEdlIePPGp4d/42vJOWh6ia",
	"//XgJ6lXWtfAFzSgES85TmWD/chQRoTTwltj7hhG6ah2JFec5ctVJQLVRHWUHzoSMJigtNzjmg5Sv6oT",
	"ZLx5liewK0D2G8jNkmQgwpA0X4OXxmMH6uPRuMECBdvSZqeMkwku9Aw97LrDYBNEP5OnCvGEIVelgaYi",
	"Ppbh33JizWvGd2WDfa2B7oZq/5l6cyYmQsU1dCmIWA5QRKPU15MMYSAN8kkiQSTKMxTnsOOMkzvKcmFA",
	"af1rhjoU96F3EJKsj+bmGOlLHiNqvHkmuEj92zjwyrCaqp3N8HN7/ACItMHSQtwJXNZhh/VaWDRFnmlG",
	"q4uLhG20+BS4ZAXqtjjmIng5TBtFzFfBIQHJzSXCMcinDDiB0leNOK6R3ggC1rlSwXIbh4WOyQLniX6U",
	"qiWfOqsvFfuD30W/jbmhrXXKA7dQodH6+9NMfZifPBeEzzPa5iXvaRHo5UyvHN61VOnXV+0Hnc1+QThh",
	"aqylKVutzlRzSyFY2MUnAx61lVFIBtSvUfEYx8Vr3BwWsEjwUjhWb3sQJZykbvQcAv3QTKy4TpmA2UMu",
	"DEttu4l+w2W+P4Os51ur+vpn34B/tknapqmQBMdT9PkZvB75gH+0zexZeH8W3uv2hajT9P1ZS/PhShzN",
	"5trHpunHsPg+8p52MJRNH2Y13h9QdzE8P/Ju/py262dl9lmZfVZmn5XZZ2X2L63MPlSL7U7I7qPGNiVD",
	"QbE7J/YjrHjYmNuwOO48PIYzl+wxw0KRcULu1FvlJt9UGDQLTA63XnrwQBn58fLyDP1wcgm8Hv5xTmLK",
	"wdenlxVoDXXMdKL3P881BjkCvWXsoNQpACrk1IXo1HMMeqBcEcrRmt0o0n1fKLThbMRPYY+7BxbLfh2l",
	"2AQ2c04SI/AsUEpI3JB+bkk64J7zKUaD7QeSEh0ienp5hjKtMxWw7c7oCmLGuB6L1oSwu+D71ZmtyVPx",
	"gINk9O787YVSTcLlheJtitc0ckPHvqeJJLxHia5yyLGexY6EggzOr4VDcpepa4MbZ5/FwSNmObcOn/BT",
	"FbBAvTVJT0a4dF8sXRVLuHk7phZcad8AYvhRq76SIR2jp13cfR+jJtZoLrsNT+7MciFMcTlji23OMQMG",
	"CHd23B19GZzODL5uPFtbCRngBU69lmD0WcnfzePamvDQUFL1olA5jYlAyXMLE5Ic0GPag0NaA5Roij5u",
	"xAsNxJeIcfRRsDSJX+iZXhqTjXhQlYY9BH/tPfLqqA5mBCWeAmqQNpZ22WV89DHZRj6hBTCsL0MOz/7g",
	"JKdopV7RdBkC9gonOF2C2oDjmBRlVKEISpP5DAfzPi9XBMWOrUBPodQvtqZSsTSxFZKsEdQqAZujeaU7",
	"zHRlGlu/gkFlUhaUMl3j0Mt9DH8fcG7NEbUA8TMkCIRB8O58ZiFQH1KmfochpDNHSPzlN9+8/tbNHWcL",
	"dDw7Ri+MMMPKUmnHs+OXXdBsxk+LZD1RtCh/VBcUNrKlYxJdoLK2JyK/5TgRKNrIKbqgy1SpPe8vlYJc",
	"1N6BiptF/Z2GTPzBK350Vvxp+IpQKTYbuqgeNUVvaXpLYgTFDAGIHct3um3KpZq3ZMrOXATKzuil1fAp",
	"Oso513UvZD2Np/xQkcsXHzfyi24h1tmc81QX+NO3+sBbU/+ymrgv55J8kg3lLGmHNQtksKKILwaS1e4n",
	"Ry9SColT/CNhSxYoPzAr4g7bwaE25cABjtWviCakL50V5deaxBXQ6xUSOWXYXdXLKeCmtMacJrHxojBO",
	"wrYa9OL8+6O//f3rb19qZVezHhhkDKda0TQhisb5CPYGfz6wS06bsvFoWOQ2vwoScRK+6Jotq9mKNEBi",
	"dm/NX8HN/qruz67l3HH14nqy2DNOMsy7qxiVUqoZEWpksYe2H2a1cpnvcDigbIXFisRNDch+hF+NUdwY",
	"ZePCAmCstYOU/oFlP/U04yAsKpt3LrXhdobdLTjD1Ttw2KAvdd209qbDS+IbiIdHTewvxa4lsbHTDn1V",
	"puAqDUqbqT6MIhaTD6N2g/EjkXoo2bLX9T0OKnTbHnvgQmMdJg8ZmhOdNMf/QlR4vs/cSXOJq2pnTd6v",
	"1GyVcTr1g9V8+l7mUiYhe5wWiou6pZD7q/0tl5dvw3UQsxxoPbjX4dA5Ozxvh0kv/gUVKY2BkqA8i9i6",
	"7r/gbYWqaub5RcI2gwhdC0LWuhJ/n7ANqLOtZpriksdNaDYuWG/DrfanuGEGz9rLpUXJxBhEdnn0epBn",
	"j+f4T/JSdryJQ5/D4JUAXENmdv8zpL7Tiech9hZTkkYaa8JK+gf10YeRcfwNAmowQ+hYU6xuoWpiIhwj",
	"XxkkAL1qBjU72b2gcWFLnj+sFPO5naerJnNDMfyyywhEVnSfZMcnXC8/rtx/G6ICtu3KPc6JyJN+4lqv",
	"3nLP1XifshpvSVY1cn20grt7r8bKaBzNdzu/1v2rBd/DbETy7dDpL8/fnSC6cIN4TaHvLZEI2xYH9tDG",
	"0XJ6Ztu06zgrMGvacIEy+lkyU861WubcBq5Vmm4UwSwvQjValVz0skftN69MQwEu9wosrMaBJp8uhbWx",
	"JMNV+jOldhepz2MgXVcM1JKcrbas1duZ2OLzrRuVdYBXs2/1v80XxatydTQsWqdB/TzSdtFWS3rLQXYC",
	"RtN74n6jUDlPZFOtl5BbxhvueiScuTowv5j8Onz6HjhbOXWtM1dT/VjfVLsmEgMvKXtHOsbpno25fMBp",
	"O/Uf2Iox0LmrNMzvLgz1OJd3m7Ub6cuBcrEKWYT6WLNysarYLMzgZlXp87JjNRXoGjfs04V4B9wGgJ/E",
	"w41HMKy3waitl7BpFZLm6xsIScSy2qqr6E9heLR1L7w7n7ktK6A8dsYMLRn7jK4r544ou10IZCgppiLi",
	"xC0QHSxUd5NLLVHIbUYjnCRbnVGUYLViAt0MuUQvyHQ5HaMbIjeEpOgbiHf726tXdqMvw1YkazAKup+q",
	"hwDTjoK2jo8PVdcr0oKY0BKVYg0AMlFUF5/kQs27IJyYhiaVOvVewF09hDkcottpEHCPOnaRo4LfTYjZ",
	"1/lnal6ZZLe6eCD0DyeNJjqbJtdujwuXTDRDm+WAWoXYcW1DDjwqZwm4c/0vZiZdovHUvR02lZW7Hg47",
	"/XVwi0sqJOFgodX1Dk84Z7yZ5ZTFF4v4ezWFyTMhanCLTgy/B+5Gl9Y/vDiazcwcEGmqLyfcGEB91R7D",
	"9GO+xumEExyDZqJnh/wC5zvLYPSqRTRHTG7y5TK8eAW++kweZnQAtT+51CZqfIXb76XZcaidnuHIrQoA",
	"TasbaCjLvBwMbQozb0QZe0PSeALeY5PI4XGntqTCIMt9d/7WbgHi4DfkBmV4SYylPtyRoMP2B4JoJNts",
	"XFYGLN5Anci4FdqiD+NRRliWFP1MqIJWIf3p5cfOI0XWmCYIxzGHztXDFJwyE6pt1yU6+DlQfoVV9fIk",
	"CdsUmVlFiLgt9ireoHq+0hjtkq407JgfN7eiqSTrF0KLKO/JDfoH2aILIlHMohzMK6a7s7bcen25Izu4",
	"DMMKN/ZVa3fioH2lbfRNFNzai5/e/+Olt8Fdtua3j+3cmpHZjBShpAsIdrFRai30kLGERtt+C8CLKHTi",
	"1srnFBmndzjaIj1deTeVXFvb/T0mWcK28AXjS5yW6TxJojuu54KIMeIEIDYGAU7JiAkTRKCMcAEh15Dv",
	"E7Z46bwGdbA2qrHEYL/XWcezggdUIFhm+YNhDEiq0P7qZOOQ4jBa8FzZ/ajeS/eqE36EU8inMn9tcAAH",
	"mMFwQm5I/LoIdPETGY7IpCzIbbuMOD2zm49S6+DXWTFAsIXcYB4ONT5EeUp/y4nTCNdiP+gT6N272fFL",
	"hIXQ0YcmXcdsKiZ3JFHvLGIc2XU0cYsV4UUqiy88GbgDTXnWBotbdiL93po0DXhSuBEVGjxHxVEb+/Ue",
	"2ha9gQP7aF9uo/gSzvLBBWhDUAfcRuE51n7idUPsbeGlKuqXh4p6F5vT9uI23E1ZSsbIC/OaK2Ws+rcb",
	"LGg0Rb+wlBSJrmoVw5v1xwK9SEHNRDjLxNjmN6l/vLQcHqdgIF/hO6gKz4kURTrim+CiYZiJBzNkSfga",
	"TKpGGShZcuVuKxxap+QqtSUHs7vOrhIrmhXqtCfomcYw3mz+B2DgF5paLdvxn9B2I22LTPwgsbqzKDrE",
	"Y5ZkVpouIfXMpFNXpfCOGMlgvfmOztvFBLqGZhwsL3pJ18DcNSK6El9J3Bss6s5mt5voZ6kalOGjQeDp",
	"n41xpWhX4CZUQjWCsiSN3aTfNIGFWErnrlorvjZeiR6rDVl6AvVovFIyBTV/VlxE/9R6Vc9q07Pa9Kw2",
	"PatNz2rTs9r0rDY9q03PatNfXm3y4p3qyU6eFtGKZ74Edd2hkA12dPQJSO3RkristvDc3jpUfyHUVLof",
	"8HuGL1wQ6U6jHZUSS7dvQL96C7+QjamhMe3os7FDIYOuOpIdxQeCQfbDSyEMaSJvyRaA5dxeJ8AffnE2",
	"iq3SabVMOnicpAR/vn5HHBJaeSEZ36k3pZCMD25MyeJwnlxrEt3Tpfg4kU1FKUEL7lY4PRDYA3oP7gL2",
	"li6AXccblnr0LouxJNUSFY3I1Pp5EdQjJM8jLVvkaoA6/dVRY0vnkjkEa+88vOKGk6jXsILfibg7oK6c",
	"rTZ27J8nsHsHR9vB3/MOr3TvHXJW4gOJe/IE27dHl6esFdlTAl1G0+lzw9rnhrWffcPaUGnZUNQ5qmD5",
	"wNJ675QiY4iii0uEa90a4u+k24fTf3fA7a4MoGe3g6ICjafxeYOcerNOOV77lhSVH8HoHxEOXMRNW9pm",
	"BGFhStdBbdoLY7v7Zvp6+hpwvVbBlskV4RsqiCaFSgqZLqk+bpj27+qbX8+/P/r2q2//dr1bRtkuMd7V",
	"Yls6fby56EDIVFgY1SqXbQYMylAJJ/F6xVLj7pqSpQBX7KFWULIbw/uSStHc1k8TadLp2guUwU9lamM9",
	"I7e9xFfzQOrE2PaPoC0ic+/Ho99yYnPjGhQ3L/3mn+rzgH5auSw9a3GwsQMgZ9PuxbXCO6AOw4CtU/59",
	"RaLbpgQk/XEwoc6xpSwwTXJOUKSmQobphIrRkeg2dM9qFJynOX63PgwCZdGaCIGXZOeybVduVlTjW1rV",
	"teEgdmfBhao31ADw3nlT1Um6ylc6N+bu7qG5x7uUnkTf4eh2g7l679YZlvSGJlRuweeEyl7YR14S7sAk",
	"5p5FHKtQLKo4uo29j/7g0pv3zagzrHZr02lbq0reVSl+30UlH6lKYwvU+hQ6bAVcHymv4Ite/rHooj7F",
	"C/oX4GpjJW35u40HGggSL++1gwNlTcmngXroQ6jX3UOQft0POnJi90nD4Y7/3uZqj0QNvg+4nyHU7t7V",
	"YHrXMtofT/Chwz8AfkOJfgC+B6i+UxeIKmUCBlXtq6ZrB+YHWWbgnrL+lKbhH/TPZMFE+OEncvPgQy0r",
	"uJJVxQPqxGWhvPf6vYzt/Y6D2fMtmNYbW9+TJPlHyjbpaUbS2bFOJz9q72jWPaaavGuacvtfGIQHGRML",
	"YhzNV0cX2sQGubyz47PdK885/fZOz74QrknMs+idtAVb3mAZrdzaR73WqxUP+ELUK2sW69q03Lfa9qGk",
	"WTXJSspMIEBVbdz5+fBfhW02Y1yOUYblCn4Cbc+xzpS47paGHjdUNogZ0XVVjBUTPmve75CGeJUaCGWz",
	"gjPvTvu5CDwUEmWZgftxvekec2o/tLTaCxXVaS794Jq4zLUxLyAAIviM2SbFa3LgVNIdm/rABEcrHbYM",
	"Wdj14CWztdKkXCtzZQ8Ud/mpd8bWp8fTTve4hU9rWY1e/Y5aLpgTmQN/R+G1XQtoWncPFIZS2xnJcDmn",
	"L5hunsTVlWurvlrMrF8n1ljHIJVehQVOvLAL52Vyd6ztaEFPT+i6uxIAHlT7rS3Io0LEurLSo/DbUJmm",
	"R0Ll8b54buuew4UERZbgba+2ox7/qbItMxEqn1ptxa9vHJoPFtZ9paTnxm7WS4h0bBBm7+2R823EDvHb",
	"+phedKrlwPD0F6/+DxDTe7mtBahS6H3slvfpL7l6Rbp2xtVfnFk+eyQNb7aHl0/fKk5Zul2zXMx13G/n",
	"BVuW7rDLQPM6G66IK03pgN3iYIc8XctFrlguFUbbbCXt1bWMt53lulHBA0RRU+rLemLP3djiVoj68eWP",
	"RxvevI9IHtpP9Hj7/NV0E7gORppTYd3zO+4WAsTnNs2uMRTe9iPFSBR9QAy1/vT+smSqdYIqMvicVgpY",
	"1IMOm+Kwh2g5mg5a0ak5+PZBd9YWBS4cuRYi8amoBYQfl7T3YZSy1NRr36HeYC9ddYhf8h5cfgtmKwsa",
	"zx/kiY3ejFYkSdh/S54LeZOwaBqTu9F4pCMzR5fqz98lLEKS4PUUWvDCIMXQ3xwc+MNqSk05HJRkw5Ed",
	"3aBQThTj96r06ZiQ918doaujyeHZzO3jqSHz9RWUH5csYm7bsgNrLXAjOvS4sptmQiNi7FvmpIcZjlZk",
	"8uX0Ve2Qm81miuHnKePLAzNWHLydHZ38cnGixkzlJ21ZqnkmXYqyxZUgEkebSHRA2OjVVC0MzhyS4oyO",
	"3oy+mr6CvaiHEVDowJzPsbAfiCJiLWPNEXXCBXkZJ6fEJmy7/43OmHACSIWJJisKfH3H4m1Rm1JTtRN4",
	"dPBRaKFay0xdElV7YNr9/b3zbsDpvnz1atDiVUdzDTNP/wFEJ/L1GvNtF6TqNDUurmPJWZ6Jg9/hv7Pj",
	"+8D9HPyu/zs7vlebW4Yyc8+J5JTcmdCvHvf1AwleV+a0zPm1oQf4D2qrJtyYqr8rHCuJ3pxk5FoAtUO0",
	"BuDSIF1/d/SJw0uI8tf+a1w/OVL0uJQ21HAYkDgwzdFL8VLHt9k4sjD9nphBwQ7O1TjfohlDHVnsPC0B",
	"y/ug885lH4HUd1zfvKB9sGC3SxiCG5kuhzwBoWqipC3Akn9PnNYhYQQxhZStEBVsi+NKbk57T69/R+A9",
	"0DM3NHvZB7b06jOzZ4zp13mjD9b0bVq0E554YRoNT7/JAi0CXB32ZeVWNxTSbc4KtQegPruJjfUaZDeh",
	"itcJY58IUq7zRNhQLVM+6P69/iC73/QE/DqPd98wXaVi/I4XX28Ltsfbry72CCiwWwO4xtiT/rhRdVgN",
	"wpBcrCqyROdrUcMRk3XsNm+CYh0gDCM32FYbpTwG5gQ7VtCiocjzvhCjo6Z0M4Z0XVNjpe4hFyUk48Ok",
	"Pki+Eg+V+boy1PZxFe1r7plbd+Ss9SHMXSA/BBdMPgSZ+HbmDnywAeqiMYkid7JGfCzokQayD0ToXHbP",
	"uNAd098HHfoDvgMJTBafOPi9yO2717/Fk0oToibrQM7r5ll4mldUcZht/erLj+23P+pPRw8E/EDTqhMR",
	"WhiTTYOgmy1a0juSIgOWHXxylbPpPN4d3mSrLHWAOJD20Wpysa0Zmywhbq7nA8wtxVall1Bu17SpFWZR",
	"+Ql+GjC/l9bfMGslM7XFkNNFGb/7Wa++TQ0GArPsYeoqwT/dO/yd5czG29csc3oH2cDCb4TZQOQ3I6ub",
	"eCuN8Pclk1WWMdn1f4BdFzaCor4idj909N50OL0gE5zGE1uTYGIVp2c8bVBBHD+4ZMjCDbSSWdBD5Hpz",
	"KERd2iYpflKWKCcrxr47f+uUTbJZmu66ajtKx/XkPAcXA9Rky0e4wX6ACZYX74u0zLoKVF8fzZ5IoKqs",
	"ao7qLN5Nie4dIzNB6Ll9fBItyJLROHomyb8QSf4VaHGQSlOhwqegPq7ziJ/proHuSpozkHKJTcfZqM9c",
	"Cozr1p6m8k/7svR0VdTat7Gno9xViBZuq9YeIv0KcRr6FmxtZNCK/dMNSZLJbco26QHLSEpdJX9SBjoX",
	"qn574rFV/u1UEAtUZ36n8LPP+mzk0GiPN9EjIWeI/n11dIFmx2eBDJzPWP2uMJHH5yEK9ZTwclAYoRpt",
	"RU1JQwbAtkS0YQpQ0lPXDi6K2lZDa93a7hWco3FU2Ne6ok+uynpNNwQJAq6GD1BZzUTLBYwKXpjnwy7p",
	"MlTgv2ldtwzoA9Y8REUGKYoJr7TbZ7GtAmB7SEI8qNpg2twbcWxK9NoMO4SXSsiSKMGy5UAsJnO3osaD",
	"TmVqU8GeN7gsrKPPqE9WLNZvS2UN1YF3Gqx2Zats69CdXBA+wUvTxcAriu6W4y58YBknd5TlItkiIiTW",
	"lZVjkwjTtKRp0uCUuvIqMGecAX0xrvMG1/jWft7YkDJMEWW98eHA0kHItl+opviOBXWR7WEIkiKW4d9y",
	"W6TNay1RdJNYY6pTAKBGj1f013qpcRqjCCfJDY5utXIRBH3RTV6WHS1MzW5zuwbSDiKoKX1s0AuUmQcX",
	"P56+e3tcKCcms//OtGmIOBNiIqgsd7tgfGlK3QQBWZQi6g3Ik1QRSVxmxjTnb0UsvSNbYXKw9N+cPhWO",
	"FV79WxfRRBtsqjqzG3UTU/RznkiaJY2LOMqapoatQicQPeZ+JEFxhd6F0VS3j2YLtLZLVYyWIdCFC4IN",
	"AqWO/v1CmPBhJVukJJI2zv3d+Vt9/+bf0FLEJrDEVETsDvJSDBUDr5OEr2lKHIB+oUCUYSjzQokA/C1K",
	"r0/R+cnR6c8/n/xyfHKsIFEkVbhCaCst2tqnWvzZkSbBabUCX3+JCT8f/guOq8ixbJVraU/jSCbpmv6b",
	"FJT0hUDkU0Y4JWlEHuF0UBZPbWw0MNYUGK9JODS17bUTyiZ9mWuzXQHIJ2nbE1QMG4RP0aGZqmxN7taQ",
	"K1utZFgIXbwNp65VBDRst5lx8eKXql4JeZOGwavBem69OrUSDDEz6KpmZpseI6uf5rJcF0ovSnwLphum",
	"2D/LbSV1WypNLZsyiZY5VlIh0RtgnC5pqn42Z6GmLRIfo4jlSay4Ak4RllJx6ob7dTe/0xU7CVW6RX7R",
	"akbnC2Cvw4A6RrWHQuj5aCmK2VERk8YTndWm/zyxfALfJMTUxvwwsincRChp18qVH0b1xNyCZULFwB8v",
	"L88u0A0UwHx3/jbcPfuD09YISm+2dAIvcuNwwgmOt7p5gCk1WrbpAkQtuy/YFkNUt8PgJia6Mk5hhf7y",
	"//7v/yNQqQGjhJW1QFol7bkG5WhIDPhXr75sUWQ/TTabzWTB+HqS84Tot9TXbMMFqcNlJkMCiO69QlJS",
	"FJttx7LAaNCITE8r6MWebBFeAFoAahtfuRKYqKRLaxvlVNyqZzQh+LahB0m4tmNRNZMuDArBhx5CKpne",
	"FMSwyOmkSNVlVTgb+YQjm/fNSUQq2k7fBgy2kGmXr+97lqdxxYoAVoOuONuyqUKhVleLZjQH41y2FZrQ",
	"dyVK0cbxtCo4sjQwuEi5V2SfZZzdlYh0ksYTKAmbZ6BCuGVlFgjrAqvoUMvxOn3O6yUGjFpPaorL1fT3",
	"p4nerKzyRFbC2qqFpXzsz7qRQUdzgaLd9ivAvJaAzgDS9UG3mUaoyMcjm0yiU9srpW91cmL4svd+z09+",
	"xU94u33vlcbZIxuIH9kcfPXls0H4P8Ug7JZzeDI2chgp5E1IvCRrku4riPQwum1lIl8HjN+3SvD5+hGx",
	"+TC6hXK7bV5W+CDEMdzCE+08I8O8+faKtq1pbDO9gmIY0sauZGv7BdRUAJzGaElkqW6+O58pTChb+oFa",
	"5Vh5sCi7PVqlQ4dweoYCO19t4XbnwVkuViR+UJLZYCG/Z3H5muntP9zsNqSHQqMrJdAE13M7vPk8HCQd",
	"22xsG7iD46O1t9Ff145VmJs+ZxtWa6/XMFX8Bzuj2kv7BNNW2v294cYSYbh2+K362j6eHVPhXjSrYKWe",
	"z8xl0NgJrqF04J/O49NuGKuGQnidSv1nNmQ+q8vPrx81BbMmxjXLy0ecYFNA8etX3wSqTOtH9hcm0aHu",
	"mg2fvv6qsZEvOkkllVt0yRh6i/mSwIAvvw0wE8bQzzjdWriLkNyuz7OLIdHY3lxZvpYzrT4Iw2pvMm9D",
	"I5HDCq83SgB8DbKMUlbzJME3CbFKabhJSXt/0dZ13E97LEfjOWimASX32JhAy3LMRql1anqBYTrTDLzg",
	"zoUfo5Tcr870ZNM+e2oU6AppI6yvQSlpxm3v42Anl6zpuHaH5TFYStSLvmYcLA+2BJVbcFv0OM99L+4R",
	"yKC+yBWnVLv+JvTz97pXTrXykpENRX6zpnX/gtVLmasIcJYvV+jq6KJKjHeZS4z2kW2OlVPEbr+C21jh",
	"NE50k2Rb7rsMP1dPiVs1RUsBTD27OUEsN0VVihi9hrIJSvE9t1vrsFc5fT3L0i1O6nFTXNXDzFfWQ9sW",
	"xbJ74aavXgUZuQFIgB07wGphvQWZtJrA3Nb9cH+6AwcoQriIwNY/W29oYSerWgH0zbiu6BUWRqlXeid4",
	"8UQOSy7ypAG5wxgCtL2/F6FFu7cOwrH1EJZudvAeOwzVluRrdHr24Z5hr2uDt7TuWHzQuvOiuUDINMG3",
	"mWRLjrOVUZU5TmO29jrjO+qtZeWkWZGygr00vrpC9uvcbVlhuLeq5RuTWhSvXs0tPbSwI4DF9dl+u+pc",
	"Q7kP3oCab9o8eXGHHUjRt1wRym3ZVQsibV2JtE+0c+/y02CQ6KX1uJA33VEATheLXghbUQccfLju/2A/",
	"kk1cMTRgUF3JR4UxvlL0HseotO3XGL5XELmd67c62rTNQxP3c+6R99pqwAgUa41Uv3+pU4zXMP2CvV8d",
	"XTSy2pB8oxfQros9OYjsIrBpvVKrw+j1flfuqfC+2ucuOn1VHZRnpzSIUFxfmAKNuNRKhI3C94DitGbD",
	"Tp5X72KnT1iAo07Rj07Qj1GY4+kqrvaN24BbPbzDFJ6/7icl6Pr9hSGDURW8/oHIQq7XCFZpE+rGEdhk",
	"VQgkaGCdIICaImkxemGGkPhle/mNH4hFYBJ7oSTPaPwEaPz4r0/4Ps/Jb/sWv5oWFlnPwJreCFynCsX1",
	"rcrkp3lXaxGWLQTDhlBo4PlsBn02gz6bQbdONYDCyumWuvALcmhvlhcMDCpn2C7qdE9sJt7f5SeoXZ9g",
	"unYEtqoUpkP/Z85IKG+8h4JysBO3oJwrJea23cUOdf67wLwk0hZdKOx4xpluLMxu3ZNpGNBdb/oxeLLL",
	"Um3hB9aUaRsYFVhc8PDSarolbreyfGwd8QUU3fp9exNOriqrobsn0JvrJdSqHaX3VUMt2AF933Uzm7pl",
	"9yqXWe2f3oML7b+a018XWYs6QTSOHJ79FLWQrs6eAlsrSw5C1id/b/thurvKIzDkPwTF/wh27DWX3yc/",
	"rvWpfxKOHOyZPYAnZz54QriqhoFBV2NY2QPrzcFBwiKcrJiQb/7r1d9fjdSFmCmqOKE91BPtBovRmsUk",
	"qQRFVfOBR3XMsvvqOU9xjIAnW8fhrQhO5ApBd/JynP6r/uP99f3/CwAA//9p29sM6D0BAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "./common.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
