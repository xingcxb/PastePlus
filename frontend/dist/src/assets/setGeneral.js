import{_ as I,r,f as T,a as u,o as V,c as B,b as e,w as a,F,u as p,i as _,g as d,p as A,h as N,d as O}from"./index.js";const G=m=>(A("data-v-0fe261bf"),m=m(),N(),m),H=G(()=>O("a",{href:"https://xingcxb.com?from=PastePlus"},"不器小窝",-1)),j={__name:"setGeneral",setup(m){let n=r(!1),b=r(!1),f=r(!1),i=r(!1),s=r("周"),c=r(35);const C={0:"天",35:"周",70:"月",100:"无限"};function g(o){switch(c.value=parseInt(o),o){case 0:s.value="天";break;case 35:s.value="周";break;case 70:s.value="月";break;case 100:s.value="无限";break}}function w(o){wails.Events.Emit({name:"handleBootUpToCore",Data:o}),wails.Events.On("handleBootUpToFrontend",l=>{l.data||(n.value=!n.value)})}function x(){wails.Events.Emit({name:"cleanAllPasteHistoryToCore"})}function S(){wails.Events.Emit({name:"loadPasteConfigToCore"}),wails.Events.On("loadPasteConfigToFrontend",o=>{switch(n.value=o.data.bootUp==="true",f.value=o.data.sound==="true",i.value=o.data.menuIcon==="true",s.value=o.data.historyCapacity,o.data.historyCapacity){case"天":c.value=0;break;case"周":c.value=35;break;case"月":c.value=70;break;case"无限":c.value=100;break}})}return T(()=>{S()}),(o,l)=>{const v=u("a-checkbox"),k=u("a-form-item"),U=u("a-slider"),E=u("a-form"),P=u("a-button"),h=u("a-col"),y=u("a-row");return V(),B(F,null,[e(E,{style:{"margin-top":"30px"},"label-col":{span:10},"wrapper-col":{span:8},layout:"horizontal",justify:"end",labelAlign:"right"},{default:a(()=>[e(k,{label:"启动"},{default:a(()=>[e(v,{checked:p(n),"onUpdate:checked":l[0]||(l[0]=t=>_(n)?n.value=t:n=t),onChange:w,class:"formStyle"},{default:a(()=>[d("开机后启动 PastePlus ")]),_:1},8,["checked"])]),_:1}),e(k,{label:"集成"},{default:a(()=>[e(v,{modelValue:p(b),"onUpdate:modelValue":l[1]||(l[1]=t=>_(b)?b.value=t:b=t),class:"formStyle"},{default:a(()=>[d("粘贴为纯文本")]),_:1},8,["modelValue"])]),_:1}),e(k,{label:"其它"},{default:a(()=>[e(v,{style:{cursor:"default"},checked:p(f),"onUpdate:checked":l[2]||(l[2]=t=>_(f)?f.value=t:f=t),disabled:"",class:"formStyle"},{default:a(()=>[d("启用音效")]),_:1},8,["checked"]),e(v,{style:{cursor:"default"},checked:p(i),"onUpdate:checked":l[3]||(l[3]=t=>_(i)?i.value=t:i=t),disabled:"",class:"formStyle"},{default:a(()=>[d("在菜单栏显示图标 ")]),_:1},8,["checked"])]),_:1}),e(k,{label:"存储时长"},{default:a(()=>[e(U,{class:"sliderStyle",marks:C,step:null,value:p(c),onChange:g,tooltipOpen:!1,"onUpdate:value":l[4]||(l[4]=t=>_(s)?s.value=t:s=t)},null,8,["value"])]),_:1})]),_:1}),e(y,null,{default:a(()=>[e(h,{offset:10},{default:a(()=>[e(P,{onClick:x},{default:a(()=>[d("清除所有记录")]),_:1})]),_:1})]),_:1}),e(y,{style:{"margin-top":"30px"}},{default:a(()=>[e(h,{offset:4},{default:a(()=>[d("友链：")]),_:1}),e(h,null,{default:a(()=>[H]),_:1})]),_:1})],64)}}},D=I(j,[["__scopeId","data-v-0fe261bf"]]);export{D as default};
