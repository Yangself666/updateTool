import{u as oe,p as k}from"./user-86fc3259.js";/* empty css                        *//* empty css                   *//* empty css                     *//* empty css                    *//* empty css                     *//* empty css                         *//* empty css                *//* empty css               *//* empty css                 */import{d as te,a as B,r as b,P as re,j as u,f as se,E as ie,S as ne,T as de,V as me,e as ue,D as pe,F as ce,G as fe,W as _e,b as ge,c as we,a2 as ve,X as he,g as M,h as D,A as t,m as f,w as a,p as _,i as o,L as ye,Y as be,o as p,a0 as w,B as d,z as R,$ as De,I as Ae,u as Fe,a1 as z,t as Ue}from"./index-83698f75.js";const ke={class:"main-wrap history-wrap"},Ve={class:"wrap-title flex-space"},Ce=D("p",null,"用户管理",-1),Ee={class:"search-area"},Pe={class:"dialog-footer"},je=te({__name:"userManage",setup(Ie){let I=B({table:[]});const N=Fe(),E=b();let c=b(!1),V=b(!1),C=b("新增用户"),A=b(!1),g=b(!1),$=/^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,r=B({dialogForm:{id:void 0,name:"",email:"",password:""},serverTypeList:[{value:1,label:"password"},{value:2,label:"secret"}],dialogFormRules:{name:[{required:!0,trigger:"blur",message:"请填写用户名称!"}],email:[{validator:(s,e,i)=>{e==""||e==null||!$.test(e)?i(new Error("请输入正确的邮箱")):i()},required:!0,trigger:"blur"}],password:[{validator:(s,e,i)=>{e==""||e&&(e==null?void 0:e.length)<6?i(new Error("请设置六位数以上的密码！")):i()},required:!0,trigger:"blur"}]},searchData:{id:void 0,name:"",email:""}});const F=oe();re(()=>{v()});function S(){C.value="新增用户",c.value=!0,A.value=!0,g.value=!1,x()}function x(){var s;(s=E.value)==null||s.clearValidate(),r.dialogForm={id:0,name:"",email:""}}function L(){var i;let s=A.value?"/api/user/add":"/api/user/edit",e=A.value?"新增成功！":"修改成功！";(i=E.value)==null||i.validate(m=>{m&&(g.value?k(F.protocolUrl+"/api/user/editPassword",{id:r.dialogForm.id,password:r.dialogForm.password}).then(n=>{n.code===200?(u({message:"密码设置成功，请告知用户！",type:"success"}),c.value=!1):u({message:n.msg})}).catch(n=>{u({message:n})}):k(F.protocolUrl+s,r.dialogForm).then(n=>{n.code===200?(u({message:e,type:"success"}),c.value=!1,v()):u({message:n.msg})}).catch(n=>{u({message:n})}))})}function Y(s){C.value="修改用户",A.value=!1,g.value=!1,c.value=!0,r.dialogForm={id:s.ID,name:s.name,email:s.email},g.value=!1}function q(s){r.dialogForm.id=s,C.value="为当前用户重新设置密码",c.value=!0,g.value=!0}function Z(){c.value=!1,x()}function j(s){z.confirm("确认删除此用户吗?若当前用户与项目已绑定，删除用户后将解除与项目的绑定！","提示",{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(()=>{k(F.protocolUrl+"/api/user/del",{id:s}).then(e=>{e.code===200?(u({message:"删除成功",type:"success"}),v()):u({message:e.msg})})}).catch(()=>{})}function v(){V.value=!0,k(F.protocolUrl+"/api/user/list",{id:Number(r.searchData.id),name:r.searchData.name,email:r.searchData.email}).then(s=>{s.code===200&&(I.table=s.data),V.value=!1}).catch(s=>{u({message:s.msg}),V.value=!1})}function H(){r.searchData={id:void 0,name:"",email:""},v()}function G(s,e){N.push({path:"/userPermManage",query:{userId:s,name:e}})}function T(s,e){let i=e==!1?"/api/user/setUserAsAdmin":"/api/user/setUserAsNonAdmin";z.confirm("该操作将改变用户的角色，是否继续操作？","提示",{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(()=>{k(F.protocolUrl+i,{id:s}).then(m=>{m.code===200?(u({message:"设置成功",type:"success"}),v()):u({message:m.msg})})}).catch(()=>{})}return(s,e)=>{const i=se,m=ie,n=ne,W=de,h=me,X=Ue("arrow-down"),J=ue,y=pe,K=ce,O=fe,Q=_e,P=ge,ee=we,ae=ve,le=he;return p(),M("div",ke,[D("div",Ve,[Ce,t(w).ApiUserAddPerm?(p(),f(i,{key:0,type:"primary",onClick:S},{default:a(()=>[d("新增")]),_:1})):_("",!0)]),D("div",Ee,[o(W,{type:"flex",justify:"end",gutter:20},{default:a(()=>[o(n,{span:5},{default:a(()=>[o(m,{modelValue:t(r).searchData.id,"onUpdate:modelValue":e[0]||(e[0]=l=>t(r).searchData.id=l),modelModifiers:{number:!0},placeholder:"用户ID",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),o(n,{span:5},{default:a(()=>[o(m,{modelValue:t(r).searchData.name,"onUpdate:modelValue":e[1]||(e[1]=l=>t(r).searchData.name=l),placeholder:"用户姓名",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),o(n,{span:5},{default:a(()=>[o(m,{modelValue:t(r).searchData.email,"onUpdate:modelValue":e[2]||(e[2]=l=>t(r).searchData.email=l),placeholder:"用户邮件",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),o(n,{span:4,class:"flex-end"},{default:a(()=>[o(i,{type:"primary",onClick:v},{default:a(()=>[d("搜索")]),_:1}),o(i,{onClick:H,type:"primary",plain:""},{default:a(()=>[d("重置")]),_:1})]),_:1})]),_:1})]),ye((p(),f(Q,{data:t(I).table,style:{width:"100%"},border:""},{default:a(()=>[o(h,{prop:"ID",label:"用户ID",align:"center",width:"90px"}),o(h,{prop:"name",label:"用户姓名",align:"center"}),o(h,{prop:"email",label:"用户邮件",align:"center"}),o(h,{prop:"isAdmin",label:"是否为管理员",align:"center"},{default:a(l=>[D("span",null,R(l.row.isAdmin===!0?"是":"否"),1)]),_:1}),o(h,{prop:"serverIdList",label:"更新时间",align:"center"},{default:a(l=>[D("span",null,R(t(De)(l.row.UpdatedAt,"YYYY-MM-DD HH:mm:ss").value),1)]),_:1}),o(h,{label:"操作",align:"center"},{default:a(l=>[o(O,null,{dropdown:a(()=>[o(K,null,{default:a(()=>[t(w).ApiUserEditPerm?(p(),f(y,{key:0,onClick:U=>Y(l.row)},{default:a(()=>[d("修改用户")]),_:2},1032,["onClick"])):_("",!0),t(w).ApiUserEditPasswordPerm?(p(),f(y,{key:1,onClick:U=>q(l.row.ID)},{default:a(()=>[d("重设密码")]),_:2},1032,["onClick"])):_("",!0),t(w).ApiUserUserPermissionPerm?(p(),f(y,{key:2,onClick:U=>G(l.row.ID,l.row.name)},{default:a(()=>[d("权限绑定")]),_:2},1032,["onClick"])):_("",!0),l.row.isAdmin===!1&&t(w).ApiUserSetUserAsAdminPerm?(p(),f(y,{key:3,onClick:U=>T(l.row.ID,l.row.isAdmin)},{default:a(()=>[d("设为管理员")]),_:2},1032,["onClick"])):_("",!0),l.row.isAdmin===!0&&t(w).ApiUserSetUserAsNonAdminPerm?(p(),f(y,{key:4,onClick:U=>T(l.row.ID,l.row.isAdmin)},{default:a(()=>[d("取消管理员")]),_:2},1032,["onClick"])):_("",!0),t(w).ApiUserDelPerm?(p(),f(y,{key:5,divided:"",onClick:U=>j(l.row.ID)},{default:a(()=>[d("删除用户")]),_:2},1032,["onClick"])):_("",!0)]),_:2},1024)]),default:a(()=>[o(i,{type:"primary"},{default:a(()=>[d(" 详情管理 "),o(J,{class:"el-icon--right"},{default:a(()=>[o(X)]),_:1})]),_:1})]),_:2},1024)]),_:1})]),_:1},8,["data"])),[[le,t(V)]]),o(ae,{modelValue:t(c),"onUpdate:modelValue":e[6]||(e[6]=l=>be(c)?c.value=l:c=l),title:t(C),width:"50%"},{footer:a(()=>[D("span",Pe,[o(i,{onClick:Z},{default:a(()=>[d("取消")]),_:1}),o(i,{type:"primary",onClick:L},{default:a(()=>[d("确定")]),_:1})])]),default:a(()=>[o(ee,{model:t(r).dialogForm,ref_key:"dialogFormRef",ref:E,"label-position":"right","label-width":"150px",size:"large",rules:t(r).dialogFormRules},{default:a(()=>[t(g)?_("",!0):(p(),M(Ae,{key:0},[o(P,{label:"用户名",prop:"name"},{default:a(()=>[o(m,{modelValue:t(r).dialogForm.name,"onUpdate:modelValue":e[3]||(e[3]=l=>t(r).dialogForm.name=l),autocomplete:"off"},null,8,["modelValue"])]),_:1}),o(P,{label:"用户email",prop:"email"},{default:a(()=>[o(m,{modelValue:t(r).dialogForm.email,"onUpdate:modelValue":e[4]||(e[4]=l=>t(r).dialogForm.email=l),autocomplete:"off"},null,8,["modelValue"])]),_:1})],64)),t(A)||t(g)?(p(),f(P,{key:1,label:"用户密码设置",prop:"password"},{default:a(()=>[o(m,{modelValue:t(r).dialogForm.password,"onUpdate:modelValue":e[5]||(e[5]=l=>t(r).dialogForm.password=l),autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),_:1})):_("",!0)]),_:1},8,["model","rules"])]),_:1},8,["modelValue","title"])])}}});export{je as default};
