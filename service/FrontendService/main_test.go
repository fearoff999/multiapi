package FrontendService

import "testing"

func Test_getItemBadge(t *testing.T) {
	type args struct {
		protected bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "item",
			args: args{
				protected: true,
			},
			want: `<span class="badge rounded-pill bg-warning"><i class="bi bi-lock"></i></span>`,
		},
		{
			name: "item",
			args: args{
				protected: false,
			},
			want: `<span class="badge rounded-pill bg-success"><i class="bi bi-unlock"></i></span>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getItemBadge(tt.args.protected); got != tt.want {
				t.Errorf("getItemBadge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateItem(t *testing.T) {
	type args struct {
		serviceName string
		protected   bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "item",
			args: args{
				serviceName: "test",
				protected:   false,
			},
			want: `
                <div id="card_test" class="card mb-3">
                    <div class="card-body">
                        <h4 class="card-title mb-0">
                            test
                            <span class="badge rounded-pill bg-success"><i class="bi bi-unlock"></i></span>
                            <a href="/test/" class="btn btn-secondary"><i class="bi bi-clipboard"></i></a>
                            <a href="/test/" class="btn btn-primary"><i class="bi bi-eye"></i></a>
                        </h4>
                    </div>
                </div>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateItem(tt.args.serviceName, tt.args.protected); got != tt.want {
				t.Errorf("generateItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateItems(t *testing.T) {
	type args struct {
		services map[string]bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "item",
			args: args{
				services: map[string]bool{
					"test": false,
				},
			},
			want: `
                <div id="card_test" class="card mb-3">
                    <div class="card-body">
                        <h4 class="card-title mb-0">
                            test
                            <span class="badge rounded-pill bg-success"><i class="bi bi-unlock"></i></span>
                            <a href="/test/" class="btn btn-secondary"><i class="bi bi-clipboard"></i></a>
                            <a href="/test/" class="btn btn-primary"><i class="bi bi-eye"></i></a>
                        </h4>
                    </div>
                </div>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateItems(tt.args.services); got != tt.want {
				t.Errorf("generateItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateHtml(t *testing.T) {
	type args struct {
		services map[string]bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "item",
			args: args{
				services: map[string]bool{
					"test": false,
				},
			},
			want: `
<html>
    <head>
        <title>Multiapi available projects</title>
        <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.4.1/font/bootstrap-icons.css">
        
        <script>
            const state = {
                APIs: [
                    'card_test',
                    
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
                        e.preventDefault();
                        copyToClipboard(e.target.href || e.target.parentNode.href);
                    });
                });
            });
        </script>

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
                
                <div id="card_test" class="card mb-3">
                    <div class="card-body">
                        <h4 class="card-title mb-0">
                            test
                            <span class="badge rounded-pill bg-success"><i class="bi bi-unlock"></i></span>
                            <a href="/test/" class="btn btn-secondary"><i class="bi bi-clipboard"></i></a>
                            <a href="/test/" class="btn btn-primary"><i class="bi bi-eye"></i></a>
                        </h4>
                    </div>
                </div>
            </div>
        </div>
    </body>
</html>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateHtml(tt.args.services); got != tt.want {
				t.Errorf("GenerateHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}
