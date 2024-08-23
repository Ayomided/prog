// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Header() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header class=\"w-full max-w-screen px-6 py-6 bg-white\"><nav class=\"container mx-auto\"><div class=\"flex flex-col sm:flex-row justify-between items-center\"><a href=\"/\" class=\"text-2xl font-bold tracking-tight mb-4 sm:mb-0 transition-colors duration-300 hover:text-gray-600\">Ji<span class=\"text-blue-500\">.</span></a><ul class=\"flex items-center space-x-6\"><li><a href=\"/articles\" class=\"text-sm uppercase tracking-wide hover:text-blue-500 transition-colors duration-300\">Blog</a></li><li><a href=\"https://github.com/Ayomided\" class=\"text-sm uppercase tracking-wide hover:text-blue-500 transition-colors duration-300\">GitHub</a></li><li><a href=\"https://www.linkedin.com/in/dadediji/\" class=\"text-sm uppercase tracking-wide hover:text-blue-500 transition-colors duration-300\">LinkedIn</a></li><li><a href=\"/feed\" class=\"text-sm uppercase tracking-wide hover:text-blue-500 transition-colors duration-300\">RSS</a></li><li><img src=\"http://localhost:8080/images/MadeByAHuman_08.svg\" target=\"_blank\" alt=\"made_by_david\"></li></ul></div></nav></header><main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
