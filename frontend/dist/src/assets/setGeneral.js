import{_ as I,r as d,f as P,a as c,o as T,c as V,b as e,w as a,F,u as f,i as _,g as r,p as A,h as N,d as O}from"./index.js";const G=m=>(A("data-v-d829a559"),m=m(),N(),m),H=G(()=>O("span",{onclick:"openBrowser('https://xingcxb.com')",class:"link"},"不器小窝",-1)),j={__name:"setGeneral",setup(m){let n=d(!1),k=d(!1),i=d(!1),p=d(!1),s=d("周"),u=d(35);const C={0:"天",35:"周",70:"月",100:"无限"};function w(o){switch(u.value=parseInt(o),o){case 0:s.value="天";break;case 35:s.value="周";break;case 70:s.value="月";break;case 100:s.value="无限";break}}function g(o){wails.Events.Emit({name:"handleBootUpToCore",Data:o}),wails.Events.On("handleBootUpToFrontend",l=>{l.data||(n.value=!n.value)})}function x(){wails.Events.Emit({name:"cleanAllPasteHistoryToCore"})}function S(){wails.Events.Emit({name:"loadPasteConfigToCore"}),wails.Events.On("loadPasteConfigToFrontend",o=>{switch(n.value=o.data.bootUp==="true",i.value=o.data.sound==="true",p.value=o.data.menuIcon==="true",s.value=o.data.historyCapacity,o.data.historyCapacity){case"天":u.value=0;break;case"周":u.value=35;break;case"月":u.value=70;break;case"无限":u.value=100;break}})}return P(()=>{S()}),(o,l)=>{const v=c("a-checkbox"),b=c("a-form-item"),U=c("a-slider"),E=c("a-form"),B=c("a-button"),h=c("a-col"),y=c("a-row");return T(),V(F,null,[e(E,{style:{"margin-top":"30px"},"label-col":{span:10},"wrapper-col":{span:8},layout:"horizontal",justify:"end",labelAlign:"right"},{default:a(()=>[e(b,{label:"启动"},{default:a(()=>[e(v,{checked:f(n),"onUpdate:checked":l[0]||(l[0]=t=>_(n)?n.value=t:n=t),onChange:g,class:"formStyle"},{default:a(()=>[r("开机后启动 PastePlus ")]),_:1},8,["checked"])]),_:1}),e(b,{label:"集成"},{default:a(()=>[e(v,{modelValue:f(k),"onUpdate:modelValue":l[1]||(l[1]=t=>_(k)?k.value=t:k=t),class:"formStyle"},{default:a(()=>[r("粘贴为纯文本")]),_:1},8,["modelValue"])]),_:1}),e(b,{label:"其它"},{default:a(()=>[e(v,{style:{cursor:"default"},checked:f(i),"onUpdate:checked":l[2]||(l[2]=t=>_(i)?i.value=t:i=t),disabled:"",class:"formStyle"},{default:a(()=>[r("启用音效")]),_:1},8,["checked"]),e(v,{style:{cursor:"default"},checked:f(p),"onUpdate:checked":l[3]||(l[3]=t=>_(p)?p.value=t:p=t),disabled:"",class:"formStyle"},{default:a(()=>[r("在菜单栏显示图标 ")]),_:1},8,["checked"])]),_:1}),e(b,{label:"存储时长"},{default:a(()=>[e(U,{class:"sliderStyle",marks:C,step:null,value:f(u),onChange:w,tooltipOpen:!1,"onUpdate:value":l[4]||(l[4]=t=>_(s)?s.value=t:s=t)},null,8,["value"])]),_:1})]),_:1}),e(y,null,{default:a(()=>[e(h,{offset:10},{default:a(()=>[e(B,{onClick:x},{default:a(()=>[r("清除所有记录")]),_:1})]),_:1})]),_:1}),e(y,{style:{"margin-top":"30px"}},{default:a(()=>[e(h,{offset:4},{default:a(()=>[r("友链：")]),_:1}),e(h,null,{default:a(()=>[H]),_:1})]),_:1})],64)}}},D=I(j,[["__scopeId","data-v-d829a559"]]);export{D as default};
