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
			want: `<div class="item__badge item__badge--protected">Protected</div>`,
		},
		{
			name: "item",
			args: args{
				protected: false,
			},
			want: `<div class="item__badge item__badge--unprotected">Unprotected</div>`,
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
                <div class="item">
                    <div class="item__name">
                        <a href="/test/">test</a>
                    </div>
                    <div class="item__badge item__badge--unprotected">Unprotected</div>
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
                <div class="item">
                    <div class="item__name">
                        <a href="/test/">test</a>
                    </div>
                    <div class="item__badge item__badge--unprotected">Unprotected</div>
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
        <style>
            
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
            }
        </style>
    </head>
    <body>
        <div class="container">
            <h1>Multiapi available projects</h1>
            <div class="items">
                
                <div class="item">
                    <div class="item__name">
                        <a href="/test/">test</a>
                    </div>
                    <div class="item__badge item__badge--unprotected">Unprotected</div>
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
