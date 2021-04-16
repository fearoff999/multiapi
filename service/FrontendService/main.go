package FrontendService

import (
	"github.com/fearoff999/multiapi/utils"
)

const scriptTpl = `
        <script>
            const state = {
                APIs: [
                    {{.ApiCardsId}}
                ],
            };
            const copyToClipboard = (text) => {
                var textArea = document.createElement("textarea");
                textArea.value = text;
                
                // Avoid scrolling to bottom
                textArea.style.top = "0";
                textArea.style.left = "0";
                textArea.style.position = "fixed";

                document.body.appendChild(textArea);
                textArea.focus();
                textArea.select();

                document.execCommand('copy');
                console.log(text, 'Copied to clipboard');

                document.body.removeChild(textArea);
            }
            const searchAPIS = (query) => {
                state.APIs.forEach(it => {
                    if (!document.getElementById(it).classList.contains('d-none')) {
                        document.getElementById(it).classList.add('d-none');
                    }
                });
                state.APIs.filter(it => (new RegExp(query, 'i')).test(it)).forEach(it => {
                    document.getElementById(it).classList.remove('d-none');
                });
            }
            window.addEventListener('DOMContentLoaded', () => {
                document.getElementById('filterApi').addEventListener('keyup', (e) => {
                    searchAPIS(e.target.value);
                });
                document.querySelectorAll('.btn-secondary').forEach(btn => {
                    btn.addEventListener('click', (e) => {
                        e.stopPropagation();
                        e.preventDefault();
                        copyToClipboard(e.target.href || e.target.parentNode.href);
                    });
                });
            });
        </script>
`

func generateScript(scriptTpl string, services map[string]bool) string {
	apiCardsId := ""
	for service := range services {
		apiCardsId += "'card_" + service + "',\n                    "
	}
	return utils.ReplaceTpl(scriptTpl, struct {
		ApiCardsId string
	}{
		ApiCardsId: apiCardsId,
	})
}

func getItemBadge(protected bool) string {
	icon := "lock"
	cls := "warning"
	if !protected {
		icon = "unlock"
		cls = "success"
	}
	badgeTpl := `<span class="badge rounded-pill bg-{{.Class}}"><i class="bi bi-{{.Icon}}"></i></span>`
	return utils.ReplaceTpl(badgeTpl, struct {
		Class string
		Icon  string
	}{
		Class: cls,
		Icon:  icon,
	})
}

func generateItem(serviceName string, protected bool) string {
	itemTpl := `
                <div id="card_{{.ServiceName}}" class="card mb-3">
                    <div class="card-body">
                        <h4 class="card-title mb-0">
                            {{.ServiceName}}
                            {{.ItemBadge}}
                            <a href="/{{.ServiceName}}/" class="btn btn-secondary"><i class="bi bi-clipboard"></i></a>
                            <a href="/{{.ServiceName}}/" class="btn btn-primary"><i class="bi bi-eye"></i></a>
                        </h4>
                    </div>
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
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.4.1/font/bootstrap-icons.css">
        {{.Script}}
        <style>
            .card {
                margin-right: 5%;
                width: 30%;
            }
            .card:nth-child(3n) {
                margin-right: 0;
            }
        </style>
    </head>
    <body class="bg-light">
        <div class="container pt-3 pb-3">
            <h1 class="mb-3 mt-3 text-primary">Multiapi available projects</h1>
            <div class="mb-3">
                <label for="filterApi" class="form-label">Name of API</label>
                <input
                    id="filterApi"
                    type="email"
                    class="form-control"
                    placeholder="some api name here..."
                >
            </div>
            <div class="d-flex flex-wrap justify-content-start">
                {{.Items}}
            </div>
        </div>
    </body>
</html>`

func GenerateHtml(services map[string]bool) string {
	return utils.ReplaceTpl(htmlSkeleton, struct {
		Script string
		Items  string
	}{
		Script: generateScript(scriptTpl, services),
		Items:  generateItems(services),
	})
}
