import{u as k,p as I}from"./user-26c67f18.js";/* empty css                         *//* empty css                     */import{d as f,o as _,g as h,h as a,z as m,A as u,i as c,w as p,B as y,u as x,j as w,C as S,D as T,F as C,G as $,q as N,s as D,_ as g,l as M,a as j,r as E,H as B,I as U,J as V,p as A,K as F,L as R,M as z,N as G,m as H,O,e as q,t as J}from"./index-91c4c1c2.js";/* empty css                */const P=l=>(N("data-v-027463f5"),l=l(),D(),l),K={class:"header-nav flex-space"},Z=P(()=>a("div",{class:"lh"},[a("p",null,"服务器资源上传工具")],-1)),Q={class:"rh flex-start"},W={class:"user-tip"},X={class:"header-info"},Y=P(()=>a("div",{class:"img-info"},[a("div",{class:"user-img"},[a("img",{src:"https://s2.loli.net/2022/04/07/gw1L2Z5sPtS8GIl.gif",alt:"avatar.gif"})]),a("p",{class:"arrow-menu icon-arrow-down"})],-1)),ee=f({__name:"headerNav",setup(l){let r=k();const d=x();function o(){I(r.protocolUrl+"/api/logout",{}).then(s=>{s.code==200&&(w({message:s.msg,type:"success"}),d.push("/login"),S())}).catch(s=>{w(s.msg)})}return(s,b)=>{const v=T,e=C,t=$;return _(),h("div",K,[Z,a("div",Q,[a("p",W,"欢迎"+m(u(r).userInfo.isAdmin?"管理员":"")+m(u(r).userInfo.name)+"登录！",1),a("div",X,[c(t,null,{dropdown:p(()=>[c(e,null,{default:p(()=>[c(v,{onClick:o},{default:p(()=>[y("退出登录")]),_:1})]),_:1})]),default:p(()=>[Y]),_:1})])])])}}});const te=g(ee,[["__scopeId","data-v-027463f5"]]),ae={key:0,class:"sidebar-container"},oe={class:"side-menu-ul"},se=["onClick"],ne=f({__name:"sideMenu",setup(l){const r=x();F();let d=M.get("userInfo").perm.menuList,o=j([{iconLabel:"UploadFilled",val:1,label:"文件上传",routePath:"/home",activeTab:1,perm_check:"upload"},{iconLabel:"Histogram",val:2,label:"历史记录",routePath:"/historyList",activeTab:2,perm_check:"history"},{iconLabel:"Management",val:3,label:"项目管理",routePath:"/resourceManage",activeTab:3,perm_check:"project",children:[{routePath:"/projectPathManage",val:31},{routePath:"/projectUserManage",val:32},{routePath:"/projectServerManage",val:33}]},{iconLabel:"Platform",val:4,label:"服务器管理",routePath:"/serverManage",activeTab:4,perm_check:"server"},{iconLabel:"Avatar",val:5,label:"用户管理",routePath:"/userManage",activeTab:5,perm_check:"user",children:[{routePath:"/userPermManage",val:51}]},{iconLabel:"Promotion",val:6,label:"权限管理",routePath:"/permissionManage",activeTab:6,perm_check:"permission"}]),s=E(0);function b(e){s.value=e.val,r.push(e.routePath)}function v(e){let t=M.get("userInfo").perm.menuList;for(let n=0;n<e.length;n++)t.indexOf(e[n].perm_check)!==-1&&(e[n].showMenu=!0);return e}return B(()=>r.currentRoute.value.path,e=>{for(let t=0;t<o.length;t++)if(o[t].routePath===e&&(s.value=o[t].val),o[t].children)for(let n=0;n<o[t].children.length;n++)o[t].children[n].routePath===e&&(s.value=o[t].val)},{immediate:!0,deep:!0}),(e,t)=>{const n=q;return u(d)&&u(d).length>0?(_(),h("div",ae,[a("ul",oe,[(_(!0),h(U,null,V(v(u(o)),(i,L)=>R((_(),h("li",{class:G(["flex-start",{"active-li":u(s)===i.activeTab}]),key:L,onClick:ue=>b(i)},[c(n,null,{default:p(()=>[(_(),H(O(i.iconLabel)))]),_:2},1024),a("p",null,m(i.label),1)],10,se)),[[z,i.showMenu&&i.showMenu==!0]])),128))])])):A("",!0)}}});const re=g(ne,[["__scopeId","data-v-e5eace6a"]]),ce={class:"index-wrap"},le={class:"main-content"},ie={class:"page-view"},_e=f({__name:"index",setup(l){return(r,d)=>{const o=J("router-view");return _(),h("div",ce,[c(te),a("div",le,[c(re),a("div",ie,[c(o)])])])}}});const fe=g(_e,[["__scopeId","data-v-3986952e"]]);export{fe as default};
