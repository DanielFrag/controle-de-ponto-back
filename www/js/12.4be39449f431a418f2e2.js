webpackJsonp([12],{172:function(e,t,r){"use strict";var i=r(2);t.a={name:"cRegisters",components:{QItem:i.j,QItemMain:i.k,QItemSeparator:i.l,QItemTile:i.m,QList:i.o},computed:{registers:function(){return this.$store.state.registers}},created:function(){this.$q.events.$on("registers-stored",this.populateView)},beforeDestroy:function(){this.$q.events.$off("registers-stored",this.populateView)},data:function(){return{data:[]}},methods:{populateView:function(){var e=this;this.data=[],this.registers.forEach(function(t){var r=new Date(t.timeStamp);e.data.push({date:r.toLocaleDateString(),time:r.toLocaleTimeString(),description:t.description,id:t._id})})},deleteConfirm:function(e){var t=this;i.a.create({title:"Confirmação",message:"Deseja realmente deletar o registro",buttons:[{label:"Cancelar",handler:function(){}},{label:"Ok",handler:function(){t.$http.delete("https://controledeponto.herokuapp.com/user/remove-register/"+e,{headers:{Authorization:"bearer "+t.$store.state.token}}).then(function(e){t.$store.commit("SET_TOKEN",e.body.token),t.data=[],i.t.create.positive({html:"Item deletado com sucesso"})},function(e){4===parseInt(e.status/100)?t.$router.replace("/"):i.t.create.negative({html:"Erro ao deletar registro"})}).finally(function(){})}}]})}}}},173:function(e,t,r){"use strict";var i=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",e._l(e.data,function(t){return r("q-list",{key:t.id,attrs:{"inset-separator":""}},[r("q-item",{on:{click:function(r){e.deleteConfirm(t.id)}}},[r("q-item-main",[r("q-item-tile",{attrs:{label:""}},[e._v("\n          "+e._s(t.time)+"\n        ")]),e._v(" "),r("q-item-tile",{attrs:{sublabel:""}},[e._v("\n          "+e._s(t.description)+"\n        ")])],1)],1)],1)}))},n=[],a={render:i,staticRenderFns:n};t.a=a},27:function(e,t,r){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var i=r(172),n=r(173),a=r(3),o=a(i.a,n.a,!1,null,null,null);t.default=o.exports}});