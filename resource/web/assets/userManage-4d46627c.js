import{u as oe,p as F}from"./user-17fe6185.js";/* empty css                        *//* empty css                   *//* empty css                     *//* empty css                    *//* empty css                     *//* empty css                         *//* empty css                *//* empty css               *//* empty css                 */import{d as te,a as B,u as re,r as h,P as se,j as c,f as ie,E as ne,S as de,T as me,V as ue,e as pe,D as ce,F as fe,G as _e,W as ge,b as we,c as ve,a2 as Ae,X as Ue,g as M,h as y,B as e,m as _,w as l,p as g,i as t,L as he,Y as ye,o as m,a0 as d,C as u,A as R,$ as De,J as be,a1 as N,t as ke}from"./index-a45fb3eb.js";const Fe={class:"main-wrap history-wrap"},Ve={class:"wrap-title flex-space"},Pe=y("p",null,"用户管理",-1),Ce={class:"search-area"},Ee={class:"dialog-footer"},je=te({__name:"userManage",setup(Ie){let I=B({table:[]});const S=re(),C=h();let f=h(!1),V=h(!1),P=h("新增用户"),D=h(!1),w=h(!1),z=/^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,r=B({dialogForm:{id:void 0,name:"",email:"",password:""},serverTypeList:[{value:1,label:"password"},{value:2,label:"secret"}],dialogFormRules:{name:[{required:!0,trigger:"blur",message:"请填写用户名称!"}],email:[{validator:(s,a,i)=>{a==""||a==null||!z.test(a)?i(new Error("请输入正确的邮箱")):i()},required:!0,trigger:"blur"}],password:[{validator:(s,a,i)=>{a==""||a&&(a==null?void 0:a.length)<6?i(new Error("请设置六位数以上的密码！")):i()},required:!0,trigger:"blur"}]},searchData:{id:void 0,name:"",email:""}});const b=oe();se(()=>{v()});function $(){P.value="新增用户",f.value=!0,D.value=!0,w.value=!1,x()}function x(){var s;(s=C.value)==null||s.clearValidate(),r.dialogForm={id:0,name:"",email:""}}function L(){var i;let s=D.value?"/api/user/add":"/api/user/edit",a=D.value?"新增成功！":"修改成功！";(i=C.value)==null||i.validate(p=>{p&&(w.value?F(b.protocolUrl+"/api/user/editPassword",{id:r.dialogForm.id,password:r.dialogForm.password}).then(n=>{n.code===200?(c({message:"密码设置成功，请告知用户！",type:"success"}),f.value=!1):c({message:n.msg})}).catch(n=>{c({message:n})}):F(b.protocolUrl+s,r.dialogForm).then(n=>{n.code===200?(c({message:a,type:"success"}),f.value=!1,v()):c({message:n.msg})}).catch(n=>{c({message:n})}))})}function Y(s){P.value="修改用户",D.value=!1,w.value=!1,f.value=!0,r.dialogForm={id:s.ID,name:s.name,email:s.email},w.value=!1}function q(s){r.dialogForm.id=s,P.value="为当前用户重新设置密码",f.value=!0,w.value=!0}function Z(){f.value=!1,x()}function j(s){N.confirm("确认删除此用户吗?若当前用户与项目已绑定，删除用户后将解除与项目的绑定！","提示",{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(()=>{F(b.protocolUrl+"/api/user/del",{id:s}).then(a=>{a.code===200?(c({message:"删除成功",type:"success"}),v()):c({message:a.msg})})}).catch(()=>{})}function v(){V.value=!0,F(b.protocolUrl+"/api/user/list",{id:Number(r.searchData.id),name:r.searchData.name,email:r.searchData.email}).then(s=>{s.code===200&&(I.table=s.data),V.value=!1}).catch(s=>{c({message:s.msg}),V.value=!1})}function H(){r.searchData={id:void 0,name:"",email:""},v()}function G(s,a){S.push({path:"/userPermManage",query:{userId:s,name:a}})}function T(s,a){let i=a==!1?"/api/user/setUserAsAdmin":"/api/user/setUserAsNonAdmin";N.confirm("该操作将改变用户的角色，是否继续操作？","提示",{confirmButtonText:"确认",cancelButtonText:"取消",type:"warning"}).then(()=>{F(b.protocolUrl+i,{id:s}).then(p=>{p.code===200?(c({message:"设置成功",type:"success"}),v()):c({message:p.msg})})}).catch(()=>{})}return(s,a)=>{const i=ie,p=ne,n=de,J=me,A=ue,W=ke("arrow-down"),X=pe,U=ce,K=fe,O=_e,Q=ge,E=we,ee=ve,ae=Ae,le=Ue;return m(),M("div",Fe,[y("div",Ve,[Pe,e(d).ApiUserAddPerm?(m(),_(i,{key:0,type:"primary",onClick:$},{default:l(()=>[u("新增")]),_:1})):g("",!0)]),y("div",Ce,[t(J,{type:"flex",justify:"end",gutter:20},{default:l(()=>[t(n,{span:5},{default:l(()=>[t(p,{modelValue:e(r).searchData.id,"onUpdate:modelValue":a[0]||(a[0]=o=>e(r).searchData.id=o),modelModifiers:{number:!0},placeholder:"用户ID",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),t(n,{span:5},{default:l(()=>[t(p,{modelValue:e(r).searchData.name,"onUpdate:modelValue":a[1]||(a[1]=o=>e(r).searchData.name=o),placeholder:"用户姓名",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),t(n,{span:5},{default:l(()=>[t(p,{modelValue:e(r).searchData.email,"onUpdate:modelValue":a[2]||(a[2]=o=>e(r).searchData.email=o),placeholder:"用户邮件",size:"large",clearable:""},null,8,["modelValue"])]),_:1}),t(n,{span:4,class:"flex-end"},{default:l(()=>[t(i,{type:"primary",onClick:v},{default:l(()=>[u("搜索")]),_:1}),t(i,{onClick:H,type:"primary",plain:""},{default:l(()=>[u("重置")]),_:1})]),_:1})]),_:1})]),he((m(),_(Q,{data:e(I).table,style:{width:"100%"},border:""},{default:l(()=>[t(A,{prop:"ID",label:"用户ID",align:"center",width:"90px"}),t(A,{prop:"name",label:"用户姓名",align:"center"}),t(A,{prop:"email",label:"用户邮件",align:"center"}),t(A,{prop:"isAdmin",label:"是否为管理员",align:"center"},{default:l(o=>[y("span",null,R(o.row.isAdmin===!0?"是":"否"),1)]),_:1}),t(A,{prop:"serverIdList",label:"更新时间",align:"center"},{default:l(o=>[y("span",null,R(e(De)(o.row.UpdatedAt,"YYYY-MM-DD HH:mm:ss").value),1)]),_:1}),t(A,{label:"操作",align:"center"},{default:l(o=>[t(O,null,{dropdown:l(()=>[t(K,null,{default:l(()=>[e(d).ApiUserEditPerm?(m(),_(U,{key:0,onClick:k=>Y(o.row)},{default:l(()=>[u("修改用户")]),_:2},1032,["onClick"])):g("",!0),e(d).ApiUserEditPasswordPerm?(m(),_(U,{key:1,onClick:k=>q(o.row.ID)},{default:l(()=>[u("重设密码")]),_:2},1032,["onClick"])):g("",!0),e(d).ApiUserUserPermissionPerm?(m(),_(U,{key:2,onClick:k=>G(o.row.ID,o.row.name)},{default:l(()=>[u("权限绑定")]),_:2},1032,["onClick"])):g("",!0),o.row.isAdmin===!1&&e(d).ApiUserSetUserAsAdminPerm?(m(),_(U,{key:3,onClick:k=>T(o.row.ID,o.row.isAdmin)},{default:l(()=>[u("设为管理员")]),_:2},1032,["onClick"])):g("",!0),o.row.isAdmin===!0&&e(d).ApiUserSetUserAsNonAdminPerm?(m(),_(U,{key:4,onClick:k=>T(o.row.ID,o.row.isAdmin)},{default:l(()=>[u("取消管理员")]),_:2},1032,["onClick"])):g("",!0),e(d).ApiUserDelPerm?(m(),_(U,{key:5,divided:"",onClick:k=>j(o.row.ID)},{default:l(()=>[u("删除用户")]),_:2},1032,["onClick"])):g("",!0)]),_:2},1024)]),default:l(()=>[e(d).ApiUserEditPerm||e(d).ApiUserEditPasswordPerm||e(d).ApiUserUserPermissionPerm||e(d).ApiUserSetUserAsAdminPerm||e(d).ApiUserSetUserAsNonAdminPerm||e(d).ApiUserDelPerm?(m(),_(i,{key:0,type:"primary"},{default:l(()=>[u(" 详情管理 "),t(X,{class:"el-icon--right"},{default:l(()=>[t(W)]),_:1})]),_:1})):g("",!0)]),_:2},1024)]),_:1})]),_:1},8,["data"])),[[le,e(V)]]),t(ae,{modelValue:e(f),"onUpdate:modelValue":a[6]||(a[6]=o=>ye(f)?f.value=o:f=o),title:e(P),width:"50%"},{footer:l(()=>[y("span",Ee,[t(i,{onClick:Z},{default:l(()=>[u("取消")]),_:1}),t(i,{type:"primary",onClick:L},{default:l(()=>[u("确定")]),_:1})])]),default:l(()=>[t(ee,{model:e(r).dialogForm,ref_key:"dialogFormRef",ref:C,"label-position":"right","label-width":"150px",size:"large",rules:e(r).dialogFormRules},{default:l(()=>[e(w)?g("",!0):(m(),M(be,{key:0},[t(E,{label:"用户名",prop:"name"},{default:l(()=>[t(p,{modelValue:e(r).dialogForm.name,"onUpdate:modelValue":a[3]||(a[3]=o=>e(r).dialogForm.name=o),autocomplete:"off"},null,8,["modelValue"])]),_:1}),t(E,{label:"用户email",prop:"email"},{default:l(()=>[t(p,{modelValue:e(r).dialogForm.email,"onUpdate:modelValue":a[4]||(a[4]=o=>e(r).dialogForm.email=o),autocomplete:"off"},null,8,["modelValue"])]),_:1})],64)),e(D)||e(w)?(m(),_(E,{key:1,label:"用户密码设置",prop:"password"},{default:l(()=>[t(p,{modelValue:e(r).dialogForm.password,"onUpdate:modelValue":a[5]||(a[5]=o=>e(r).dialogForm.password=o),autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),_:1})):g("",!0)]),_:1},8,["model","rules"])]),_:1},8,["modelValue","title"])])}}});export{je as default};