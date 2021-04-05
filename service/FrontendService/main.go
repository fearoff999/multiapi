package FrontendService

import "github.com/fearoff999/multiapi/utils"

const styles = `
            .container {
                display: flex;
                flex-direction: column;
                max-width: 1024px;
                margin: auto;
            }

            .items {
                display: flex;
                flex-direction: column;
            }

            .item {
                display: flex;
                background: rgb(243, 240, 214);
                color: rgb(28, 54, 70);
                padding: 5px 30px;
                margin-bottom: 15px;
                border-radius: 10px;
                font-weight: 700;
                font-family: "Roboto Slab", "Times New Roman", serif;
                font-size: 1em;
                line-height: 1.55em;
                align-items: center;
                justify-content: space-around;
            }

            .item:nth-child(even) {
                background: rgb(99, 106, 94);
                color: #FFFFFF;
            }
            .item a {
                color: rgb(28, 54, 70);
            }
            .item:nth-child(even) a {
                color: #FFFFFF;
            }
            .item__name {
                margin-right: 15px;
                flex-basis: 70%;
            }
            .item__badge {
                border-radius: 2px;
                border: 1px solid;
                padding: 5px 10px;
                border-radius: 15px;
                color: #FFFFFF;
            }
            .item__badge--protected {
                border-color: rgb(76, 180, 130);
                background: light-green;
                background: rgb(76, 180, 130);
            }
            .item__badge--unprotected {
                border-color: rgb(202, 163, 74);
                background: rgb(202, 163, 74);
            }`

func getItemBadge(protected bool) string {
	text := "Protected"
	cls := "item__badge--protected"
	if !protected {
		text = "Unprotected"
		cls = "item__badge--unprotected"
	}
	badgeTpl := `<div class="item__badge {{.Class}}">{{.Text}}</div>`
	return utils.ReplaceTpl(badgeTpl, struct {
		Class string
		Text  string
	}{
		Class: cls,
		Text:  text,
	})
}

func generateItem(serviceName string, protected bool) string {
	itemTpl := `
                <div class="item">
                    <div class="item__name">
                        <a href="/{{.ServiceName}}/">{{.ServiceName}}</a>
                    </div>
                    {{.ItemBadge}}
                </div>`
	return utils.ReplaceTpl(itemTpl, struct {
		ServiceName string
		ItemBadge   string
	}{
		ServiceName: serviceName,
		ItemBadge:   getItemBadge(protected),
	})
}

func generateItems(services map[string]bool) string {
	res := ""

	for serviceName, protected := range services {
		res += generateItem(serviceName, protected)
	}

	return res
}

const htmlSkeleton = `
<html>
    <head>
        <title>Multiapi available projects</title>
        <style>
            {{.CSS}}
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Multiapi available projects</h1>
            <div class="items">
                {{.Items}}
            </div>
        </div>
    </body>
</html>`

func GenerateHtml(services map[string]bool) string {
	return utils.ReplaceTpl(htmlSkeleton, struct {
		CSS   string
		Items string
	}{
		CSS:   styles,
		Items: generateItems(services),
	})
}
