webpackJsonp([13],{174:function(e,t,a){"use strict";var o=a(2);t.a={name:"cInsert",components:{QBtn:o.c,QIcon:o.g,QInlineDatetime:o.h,QInput:o.i,QSelect:o.q},data:function(){return{date:new Date,obs:"",options:[{label:"",value:""},{label:"Entrada",value:"Entrada"},{label:"Saída",value:"Saída"}],select:""}},methods:{submitConfirm:function(){var e=this;o.a.create({title:"Confirmação",message:"Deseja realmente enviar o registro",buttons:[{label:"Cancelar",handler:function(){}},{label:"Ok",handler:function(){e.date.setSeconds(0),e.date.setMilliseconds(0);var t="";e.select&&(t=""+e.select),e.obs&&(t=t+": "+e.obs),e.$http.post("https://controledeponto.herokuapp.com/user/register-date",{timeStamp:e.date,description:t},{headers:{Authorization:"bearer "+e.$store.state.token}}).then(function(t){e.$store.commit("SET_TOKEN",t.body.token),o.t.create.positive({html:"Registro enviado"})},function(t){4===parseInt(t.status/100)?e.$router.replace("/"):o.t.create.negative({html:"Erro ao criar registro"})}).finally(function(){})}}]})}}}},175:function(e,t,a){"use strict";var o=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"layout-padding docs-input row justify-center"},[a("form",[a("q-inline-datetime",{attrs:{type:"datetime",color:"primary",format24h:""},model:{value:e.date,callback:function(t){e.date=t},expression:"date"}}),e._v(" "),a("q-select",{attrs:{"stack-label":"Descrição",inverted:"",color:"secondary",separator:"",options:e.options},model:{value:e.select,callback:function(t){e.select=t},expression:"select"}}),e._v(" "),a("q-input",{attrs:{type:"textarea",color:"secondary",inverted:"","min-rows":5,placeholder:"Observação"},model:{value:e.obs,callback:function(t){e.obs=t},expression:"obs"}}),e._v(" "),a("br"),e._v(" "),a("br"),e._v(" "),a("br"),e._v(" "),a("q-btn",{staticClass:"fixed",staticStyle:{right:"18px",bottom:"18px"},attrs:{round:"",color:"primary",glossy:""},on:{click:e.submitConfirm}},[a("q-icon",{attrs:{name:"add"}})],1)],1)])},n=[],r={render:o,staticRenderFns:n};t.a=r},28:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var o=a(174),n=a(175),r=a(3),s=r(o.a,n.a,!1,null,null,null);t.default=s.exports}});