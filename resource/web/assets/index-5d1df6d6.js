import{u as M}from"./user-4e4394e2.js";/* empty css                         *//* empty css                     */import{d as v,o as _,g as d,h as t,y as h,z as p,i as r,w as u,A as P,u as b,B as k,C as L,D as I,F as y,p as S,q as T,_ as m,a as C,r as $,G as D,H as N,I as j,J as B,l as E,K as U,L as A,M as F,k as R,N as V,e as z,s as G}from"./index-6873b35d.js";/* empty css                */const w=c=>(S("data-v-800e4d73"),c=c(),T(),c),H={class:"header-nav flex-space"},q=w(()=>t("div",{class:"lh"},[t("p",null,"服务器资源上传工具")],-1)),J={class:"rh flex-start"},K={class:"user-tip"},O={class:"header-info"},Z=w(()=>t("div",{class:"img-info"},[t("div",{class:"user-img"},[t("img",{src:"https://s2.loli.net/2022/04/07/gw1L2Z5sPtS8GIl.gif",alt:"avatar.gif"})]),t("p",{class:"arrow-menu icon-arrow-down"})],-1)),Q=v({__name:"headerNav",setup(c){let l=M();const s=b();function n(){s.push("/login"),k()}return(f,g)=>{const e=L,a=I,o=y;return _(),d("div",H,[q,t("div",J,[t("p",K,"欢迎"+h(p(l).userInfo.isAdmin?"管理员":"")+h(p(l).userInfo.name)+"登录！",1),t("div",O,[r(o,null,{dropdown:u(()=>[r(a,null,{default:u(()=>[r(e,{onClick:n},{default:u(()=>[P("退出登录")]),_:1})]),_:1})]),default:u(()=>[Z]),_:1})])])])}}});const W=m(Q,[["__scopeId","data-v-800e4d73"]]),X={class:"sidebar-container"},Y={class:"side-menu-ul"},ee=["onClick"],ae=v({__name:"sideMenu",setup(c){const l=b();B();let s=C([{iconLabel:"UploadFilled",val:1,label:"文件上传",routePath:"/home",activeTab:1,perm_check:"upload"},{iconLabel:"Histogram",val:2,label:"历史记录",routePath:"/historyList",activeTab:2,perm_check:"history"},{iconLabel:"Management",val:3,label:"项目管理",routePath:"/resourceManage",activeTab:3,perm_check:"project",children:[{routePath:"/projectPathManage",val:31},{routePath:"/projectUserManage",val:32},{routePath:"/projectServerManage",val:33}]},{iconLabel:"Platform",val:4,label:"服务器管理",routePath:"/serverManage",activeTab:4,perm_check:"server"},{iconLabel:"Avatar",val:5,label:"用户管理",routePath:"/userManage",activeTab:5,perm_check:"user",children:[{routePath:"/userPermManage",val:51}]},{iconLabel:"Promotion",val:6,label:"权限管理",routePath:"/permissionManage",activeTab:6,perm_check:"permission"}]),n=$(0);function f(e){n.value=e.val,l.push(e.routePath)}console.log("currentSelecTab====",n.value);function g(e){let a=E.get("userInfo").perm.menuList;for(let o=0;o<e.length;o++)a.indexOf(e[o].perm_check)!==-1&&(e[o].showMenu=!0);return e}return D(()=>l.currentRoute.value.path,e=>{console.log("toPath----",e);for(let a=0;a<s.length;a++)if(s[a].routePath===e&&(n.value=s[a].val),s[a].children)for(let o=0;o<s[a].children.length;o++)s[a].children[o].routePath===e&&(n.value=s[a].val)},{immediate:!0,deep:!0}),(e,a)=>{const o=z;return _(),d("div",X,[t("ul",Y,[(_(!0),d(N,null,j(g(p(s)),(i,x)=>U((_(),d("li",{class:F(["flex-start",{"active-li":p(n)===i.activeTab}]),key:x,onClick:ce=>f(i)},[r(o,null,{default:u(()=>[(_(),R(V(i.iconLabel)))]),_:2},1024),t("p",null,h(i.label),1)],10,ee)),[[A,i.showMenu&&i.showMenu==!0]])),128))])])}}});const te=m(ae,[["__scopeId","data-v-fa26160e"]]),oe={class:"index-wrap"},se={class:"main-content"},ne={class:"page-view"},re=v({__name:"index",setup(c){return(l,s)=>{const n=G("router-view");return _(),d("div",oe,[r(W),t("div",se,[r(te),t("div",ne,[r(n)])])])}}});const pe=m(re,[["__scopeId","data-v-3986952e"]]);export{pe as default};
