import{u as Q,p as y}from"./user-4e4394e2.js";/* empty css                        *//* empty css                   *//* empty css                     *//* empty css                    *//* empty css                     *//* empty css                         *//* empty css                *//* empty css               *//* empty css                 */import{d as ee,a as M,r as D,O as te,j as m,f as oe,E as ae,R as le,S as re,U as ne,e as se,C as ie,D as ce,F as de,V as pe,b as me,c as ue,a1 as fe,W as _e,g as ge,h as b,z as r,k as f,w as e,m as _,i as o,K as je,X as ve,o as u,$ as g,A as i,y as he,Z as we,u as ye,a0 as A,s as De}from"./index-6873b35d.js";const be={class:"main-wrap history-wrap"},ke={class:"wrap-title flex-space"},Ie=b("p",null,"项目管理",-1),Ce={class:"search-area"},Ne={class:"dialog-footer"},Se=ee({__name:"resourceManage",setup(Fe){const T=ye();let N=M({table:[],serverList:[]});const F=D();let c=D(!1),k=D(!1),V=D("新增项目"),I=D(!1),n=M({dialogForm:{projectName:"",projectIntro:"",id:void 0},dialogFormRules:{projectName:[{required:!0,trigger:"blur",message:"请填写项目名称!"}],serverIdList:[{required:!0,trigger:"blur",message:"请选择绑定的服务器!"}]},searchData:{id:void 0,projectName:""}});const j=Q();te(()=>{L(),v()});function B(){c.value=!0,I.value=!0,V.value="新增项目",P()}function L(){y(j.protocolUrl+"/api/server/list",{}).then(t=>{t.code===200&&(N.serverList=t.data)})}function R(){var s;let t=I.value?"/api/project/add":"/api/project/edit",a=I.value?"新增成功！":"修改成功！";(s=F.value)==null||s.validate(d=>{d&&y(j.protocolUrl+t,n.dialogForm).then(p=>{p.code===200?(m({message:a,type:"success"}),c.value=!1,v()):m({message:p.msg})}).catch(p=>{m({message:p})})})}function v(){k.value=!0,y(j.protocolUrl+"/api/project/list",{id:Number(n.searchData.id),projectName:n.searchData.projectName}).then(t=>{t.code===200?N.table=t.data:m({message:t.msg}),k.value=!1}).catch(t=>{m({message:t.msg}),k.value=!1})}function S(t){V.value="修改项目",I.value=!1,c.value=!0,n.dialogForm={projectName:t.projectName,projectIntro:t.projectIntro,id:t.ID}}function P(){var t;n.dialogForm={projectName:"",projectIntro:"",id:void 0},(t=F.value)==null||t.clearValidate()}function z(){c.value=!1,P()}function $(t){y(j.protocolUrl+"/api/project/delCheck",{id:t}).then(a=>{a.code===200?a.data=="true"||a.data==!0?A.confirm("该项目下存在关联路径，关联服务器以及关联用户，是否强制删除？若强制删除则该项目下的信息无法恢复！","提示",{confirmButtonText:"强制删除",cancelButtonText:"取消",type:"warning"}).then(()=>{U(t)}).catch(()=>{}):A.confirm("该项目删除后的信息无法恢复，是否删除？","提示",{confirmButtonText:"删除",cancelButtonText:"取消",type:"warning"}).then(()=>{U(t)}).catch(()=>{}):m({message:a.msg})})}function U(t){y(j.protocolUrl+"/api/project/del",{id:t}).then(a=>{a.code===200?(m({message:"删除成功",type:"success"}),v()):m({message:a.msg})})}function Y(){n.searchData={id:void 0,projectName:""},v()}function E(t,a,s){let d=s===1?"/projectUserManage":s===2?"/projectPathManage":"/projectServerManage";T.push({path:d,query:{projectId:t,projectName:a}})}return(t,a)=>{const s=oe,d=ae,p=le,q=re,h=ne,H=De("arrow-down"),K=se,w=ie,O=ce,W=de,X=pe,x=me,Z=ue,G=fe,J=_e;return u(),ge("div",be,[b("div",ke,[Ie,r(g).ApiProjectAddPerm?(u(),f(s,{key:0,type:"primary",onClick:B},{default:e(()=>[i("新增")]),_:1})):_("",!0)]),b("div",Ce,[o(q,{gutter:20,type:"flex",justify:"end"},{default:e(()=>[o(p,{span:5},{default:e(()=>[o(d,{modelValue:r(n).searchData.id,"onUpdate:modelValue":a[0]||(a[0]=l=>r(n).searchData.id=l),modelModifiers:{number:!0},placeholder:"项目ID",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),o(p,{span:5},{default:e(()=>[o(d,{modelValue:r(n).searchData.projectName,"onUpdate:modelValue":a[1]||(a[1]=l=>r(n).searchData.projectName=l),placeholder:"项目名称",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),o(p,{span:4,class:"flex-end"},{default:e(()=>[o(s,{type:"primary",onClick:v},{default:e(()=>[i("搜索")]),_:1}),o(s,{onClick:Y,type:"primary",plain:""},{default:e(()=>[i("重置")]),_:1})]),_:1})]),_:1})]),je((u(),f(X,{data:r(N).table,style:{width:"100%"},border:""},{default:e(()=>[o(h,{prop:"ID",label:"项目ID",align:"center"}),o(h,{prop:"projectName",label:"项目名称",align:"center"}),o(h,{prop:"projectIntro",label:"项目简介",align:"center"}),o(h,{prop:"serverIdList",label:"更新时间",align:"center"},{default:e(l=>[b("span",null,he(r(we)(l.row.UpdatedAt,"YYYY-MM-DD HH:mm:ss").value),1)]),_:1}),o(h,{label:"操作",align:"center"},{default:e(l=>[o(W,null,{dropdown:e(()=>[o(O,null,{default:e(()=>[r(g).ApiProjectEditPerm?(u(),f(w,{key:0,onClick:C=>S(l.row)},{default:e(()=>[i("修改项目")]),_:2},1032,["onClick"])):_("",!0),r(g).ApiProjectUserListPerm?(u(),f(w,{key:1,onClick:C=>E(l.row.ID,l.row.projectName,1)},{default:e(()=>[i("用户绑定")]),_:2},1032,["onClick"])):_("",!0),r(g).ApiProjectServerListPerm?(u(),f(w,{key:2,onClick:C=>E(l.row.ID,l.row.projectName,3)},{default:e(()=>[i("服务器绑定")]),_:2},1032,["onClick"])):_("",!0),r(g).ApiProjectPathPerm?(u(),f(w,{key:3,onClick:C=>E(l.row.ID,l.row.projectName,2)},{default:e(()=>[i("路径管理")]),_:2},1032,["onClick"])):_("",!0),r(g).ApiProjectDelPerm?(u(),f(w,{key:4,divided:"",onClick:C=>$(l.row.ID)},{default:e(()=>[i("删除项目")]),_:2},1032,["onClick"])):_("",!0)]),_:2},1024)]),default:e(()=>[o(s,{type:"primary"},{default:e(()=>[i(" 详情管理 "),o(K,{class:"el-icon--right"},{default:e(()=>[o(H)]),_:1})]),_:1})]),_:2},1024)]),_:1})]),_:1},8,["data"])),[[J,r(k)]]),o(G,{modelValue:r(c),"onUpdate:modelValue":a[4]||(a[4]=l=>ve(c)?c.value=l:c=l),title:r(V),width:"50%"},{footer:e(()=>[b("span",Ne,[o(s,{onClick:z},{default:e(()=>[i("取消")]),_:1}),o(s,{type:"primary",onClick:R},{default:e(()=>[i(" 确定 ")]),_:1})])]),default:e(()=>[o(Z,{model:r(n).dialogForm,ref_key:"dialogFormRef",ref:F,"label-position":"right","label-width":"150px",size:"large",rules:r(n).dialogFormRules},{default:e(()=>[o(x,{label:"项目名称",prop:"projectName"},{default:e(()=>[o(d,{modelValue:r(n).dialogForm.projectName,"onUpdate:modelValue":a[2]||(a[2]=l=>r(n).dialogForm.projectName=l),autocomplete:"off"},null,8,["modelValue"])]),_:1}),o(x,{label:"项目简介",prop:"projectIntro"},{default:e(()=>[o(d,{modelValue:r(n).dialogForm.projectIntro,"onUpdate:modelValue":a[3]||(a[3]=l=>r(n).dialogForm.projectIntro=l),autocomplete:"off",rows:3,type:"textarea",resize:"none"},null,8,["modelValue"])]),_:1})]),_:1},8,["model","rules"])]),_:1},8,["modelValue","title"])])}}});export{Se as default};
