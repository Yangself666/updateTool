import{u as x,p}from"./user-86fc3259.js";/* empty css                          *//* empty css                    */import{d as U,a as B,r as b,P as D,a3 as E,a4 as M,f as w,g as k,h as m,z as _,A as o,i as h,w as l,m as S,p as P,K as A,j as q,o as c,B as u,I as R,J as T,a0 as z,u as F}from"./index-83698f75.js";const G={class:"main-wrap project-user-wrap"},J={class:"wrap-title"},K={class:"bottom-btn flex-end"},X=U({__name:"projectServerManage",setup($){const d=A(),v=F(),i=x();let e=B({table:[],checkedServer:[],userIdList:[],tempServerList:[]});const s=b(!1),n=b(!0);D(()=>{I(),C()});function I(){p(i.protocolUrl+"/api/server/list",{}).then(r=>{r.code===200&&(e.table=r.data)})}function C(){p(i.protocolUrl+"/api/project/server/list",{id:Number(d.query.projectId)}).then(r=>{if(r.code===200&&r.data.length>0){for(let t=0;t<r.data.length;t++)e.tempServerList.push(r.data[t].serverName),e.userIdList.push(r.data[t].ID);e.checkedServer=e.userIdList,e.checkedServer.length==e.table.length?(n.value=!1,s.value=!0):s.value=!1}})}const L=r=>{if(e.table.length>0&&r==!0)for(let t=0;t<e.table.length;t++)e.tempServerList.push(e.table[t].serverName),e.userIdList.push(e.table[t].ID);else e.userIdList=[];e.checkedServer=e.userIdList,n.value=!1},j=r=>{const t=r.length;s.value=t===e.table.length,n.value=t>0&&t<e.table.length};function y(){p(i.protocolUrl+"/api/project/server/edit",{projectId:Number(d.query.projectId),serverIdList:e.checkedServer}).then(r=>{r.code==200&&(q({message:"绑定成功",type:"success"}),v.push("/resourceManage"))})}function N(){v.push("/resourceManage")}return(r,t)=>{const f=E,V=M,g=w;return c(),k("div",G,[m("div",J,[m("p",null,_(o(d).query.projectName)+"-服务器绑定管理",1)]),h(f,{modelValue:s.value,"onUpdate:modelValue":t[0]||(t[0]=a=>s.value=a),indeterminate:n.value,onChange:t[1]||(t[1]=a=>L(s.value))},{default:l(()=>[u(" 全部服务器 ")]),_:1},8,["modelValue","indeterminate"]),h(V,{modelValue:o(e).checkedServer,"onUpdate:modelValue":t[2]||(t[2]=a=>o(e).checkedServer=a),onChange:j},{default:l(()=>[(c(!0),k(R,null,T(o(e).table,a=>(c(),S(f,{key:a.ID,label:a.ID},{default:l(()=>[u(_(a.serverName),1)]),_:2},1032,["label"]))),128))]),_:1},8,["modelValue"]),m("div",K,[h(g,{type:"primary",onClick:N,class:"mt35"},{default:l(()=>[u(" 返回 ")]),_:1}),o(z).ApiProjectServerEditPerm?(c(),S(g,{key:0,type:"primary",onClick:y,class:"mt35"},{default:l(()=>[u(" 确定修改 ")]),_:1})):P("",!0)])])}}});export{X as default};