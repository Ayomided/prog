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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex items-center justify-between flex-wrap py-8\"><div class=\"lg:flex-grow\"><div class=\"flex gap-4\"><a href=\"/\" class=\"block lg:inline-block font-semibold rounded border px-3 py-1 tracking-wide text-sm hover:text-blue-500 transition-colors duration-300\">Home</a> <a href=\"/about\" class=\"text-sm block lg:inline-block font-semibold rounded border px-3 py-1 tracking-wide hover:text-blue-500 transition-colors duration-300\">About</a> <a href=\"https://github.com/Ayomided\" class=\"text-sm block lg:inline-block font-semibold rounded border px-3 py-1 tracking-wide hover:text-blue-500 transition-colors duration-300\">GitHub</a> <a href=\"https://www.linkedin.com/in/dadediji/\" class=\"text-sm block lg:inline-block font-semibold rounded border px-3 py-1 tracking-wide hover:text-blue-500 transition-colors duration-300\">LinkedIn</a> <a href=\"/feed\" class=\"text-sm block lg:inline-block font-semibold rounded border px-3 py-1 tracking-wide hover:text-blue-500 transition-colors duration-300\">RSS</a></div></div></div><main>")
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
