package css

// TODO: Color scheme should be passed into this function from the template
// backend, and it will be defined in the config or some part of the
// application that is accessible to the web application developer.
func Overrides() string {
	return `nav.navbar h1.title{margin-top:6px;padding-right:35px;padding-left:12px;}
div.sidebar{text-align:center;}	
a.side-menu-item{background:#f5f5f5;font-size:14px;color:#808080;padding:15px;width:100%;}
.sidebar-item{background:#f5f5f5;text-align:center;padding-top:18px;font-size:18px;font-weight:500;}
.sidebar-menu{text-align:center;padding-top:18px;font-size:18px;font-weight:500;}
.sidebar-header{text-align:left;padding-top:18px;font-size:24px;font-weight:500;}
.sidebar-li{margin-bottom:12px;margin-top:12px;text-align:left;font-size:20px;font-weight:300;}
.register-button{width:100%;}
.top-spacer{margin-top:45px;}
.subhead-text{font-weight:300;}
.forgot-password{margin-top:10px;text-align:right;}
.forgot-password a{color:#c0c0c0;}
.forgot-password a:hover{color:gray}
a.site-title{color:#363636;}
.footer-title{font-weight:800;font-size:16px;}
.footer-description{margin-bottom:14px;font-weight:300;}
.footer-li{padding:5px;}
`
}
