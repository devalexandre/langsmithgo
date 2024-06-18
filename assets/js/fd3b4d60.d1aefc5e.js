"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[383],{8635:(t,n,e)=>{e.r(n),e.d(n,{assets:()=>r,contentTitle:()=>o,default:()=>l,frontMatter:()=>i,metadata:()=>u,toc:()=>c});var a=e(4848),s=e(8453);const i={sidebar_position:3,sidebar_label:"Patching"},o=void 0,u={id:"usage/patching",title:"patching",description:"Patching an Existing Run",source:"@site/docs/usage/patching.md",sourceDirName:"usage",slug:"/usage/patching",permalink:"/langsmithgo/docs/usage/patching",draft:!1,unlisted:!1,editUrl:"https://github.com/devalexandre/langsmithgo/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/usage/patching.md",tags:[],version:"current",sidebarPosition:3,frontMatter:{sidebar_position:3,sidebar_label:"Patching"},sidebar:"tutorialSidebar",previous:{title:"Posting",permalink:"/langsmithgo/docs/usage/posting"},next:{title:"Single Request",permalink:"/langsmithgo/docs/usage/single-request"}},r={},c=[{value:"Patching an Existing Run",id:"patching-an-existing-run",level:2}];function d(t){const n={code:"code",h2:"h2",p:"p",pre:"pre",...(0,s.R)(),...t.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(n.h2,{id:"patching-an-existing-run",children:"Patching an Existing Run"}),"\n",(0,a.jsxs)(n.p,{children:["The ",(0,a.jsx)(n.code,{children:"PatchRun"})," function allows updating an existing run, perhaps to add results, change status, or log additional data as the run progresses. For instance, if a run is completed and you want to update it with output data or change its status, you would use PATCH to modify only those specific fields without touching the rest of the run's data."]}),"\n",(0,a.jsx)(n.pre,{children:(0,a.jsx)(n.code,{className:"language-go",children:'patchInput := &langsmithgo.RunPayload{\n    RunID:    "unique-run-id",\n    Outputs:  map[string]interface{}{"output_key": "output_value"},\n    Events:   []langsmithgo.Event{{EventName: "event1", Value: "value1"}},\n}\n\nerr := client.PatchRun("unique-run-id", patchInput)\nif err != nil {\n    log.Fatalf("Failed to patch run: %v", err)\n}\n'})})]})}function l(t={}){const{wrapper:n}={...(0,s.R)(),...t.components};return n?(0,a.jsx)(n,{...t,children:(0,a.jsx)(d,{...t})}):d(t)}},8453:(t,n,e)=>{e.d(n,{R:()=>o,x:()=>u});var a=e(6540);const s={},i=a.createContext(s);function o(t){const n=a.useContext(i);return a.useMemo((function(){return"function"==typeof t?t(n):{...n,...t}}),[n,t])}function u(t){let n;return n=t.disableParentContext?"function"==typeof t.components?t.components(s):t.components||s:o(t.components),a.createElement(i.Provider,{value:n},t.children)}}}]);