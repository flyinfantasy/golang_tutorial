"use strict";(self.webpackChunkx_race=self.webpackChunkx_race||[]).push([[647],{4647:function(e,n,t){t.r(n);var r=t(3433),o=t(1413),a=t(9439),i=t(2791),u=t(1632),s=t(9806),c=t(6937),m=t(961),d=t(3465),l=t(6251),_=t(9240),g=t(184);n.default=function(){var e=(0,i.useContext)(d.U),n=e.$,t=e.setReducer,p=e.getReducer,f=(0,i.useRef)(null),x=p((function(e){var n=e.dashBoardChildCurrentBack,t=e.dashBoard,r=e.login;return{paginationObj:n.paginationObj,paginationOptions:n.paginationOptions,pigeonInductionData:n.pigeonInductionData,choiceDate:t.choiceDate,choicePigeonLoftCode:t.choicePigeonLoftCode,clubCode:r.loginInfo.fromListClub||r.loginInfo.club,clubDate:r.loginInfo.currentDate,reflashCount:t.reflashCount}})),C=x.paginationObj,I=x.paginationOptions,b=x.pigeonInductionData,h=x.choiceDate,w=x.choicePigeonLoftCode,D=x.clubCode,N=x.clubDate,v=x.reflashCount,j=(0,i.useState)({currentRenderItem:[],sortColumn:"",reverseSort:!1}),P=(0,a.Z)(j,2),Q=P[0],R=Q.currentRenderItem,O=Q.sortColumn,Z=Q.reverseSort,k=P[1],B=(0,o.Z)((0,o.Z)({gridHeaderItem:{gridHeaderWidth:["minmax(85px, 5%)","minmax(85px, 8.33%)","minmax(105px, 8.33%)","minmax(110px, 8.33%)","minmax(156px, 12.5%)","minmax(156px, 12.5%)","minmax(156px, 12.5%)","minmax(202px, 16.25%)","minmax(202px, 16.25%)"],renderHeaderItem:n.maps([{sortColumn:"rowNum",columnName:"\u7b46\u6578"},{sortColumn:"pigeon_number",columnName:"\u9d3f\u820d\u7de8\u865f"},{sortColumn:"",columnName:"\u74b0\u865f"},{sortColumn:"",columnName:"\u92c1\u7247\u865f\u78bc"},{sortColumn:"reset_get_web_Induction_date",columnName:"\u9d3f\u9418\u6642\u9593"},{sortColumn:"reset_error_range_time",columnName:"\u8aa4\u5dee\u6642\u9593"},{sortColumn:"reset_induction_date",columnName:"\u638c\u4e2d\u5bf6\u6642\u9593"},{sortColumn:"",columnName:"UID"},{sortColumn:"",columnName:"\u7d93\u7def\u5ea6"}],(function(e,n){var t=e.sortColumn,r=e.columnName;return(0,g.jsxs)("div",{className:t?"use-sort":"",onClick:function(){return k((function(e){return(0,o.Z)((0,o.Z)({},e),{},{sortColumn:t,reverseSort:!Z})}))},children:[r,t?(0,g.jsx)(s.G,{icon:u.CmM}):""]},n)}))},gridBodyItem:{renderBodyItem:R.length>0?n.maps(R,(function(e,n){return(0,g.jsxs)("div",{className:n%2!==0?"table-grid-body-column bg-double":"table-grid-body-column bg-single",style:{gridTemplateColumns:"minmax(85px, 5%) minmax(85px, 8.33%) minmax(105px, 8.33%) minmax(110px, 8.33%) minmax(156px, 12.5%) minmax(156px, 12.5%) minmax(156px, 12.5%) minmax(202px, 16.25%) minmax(202px, 16.25%)"},children:[(0,g.jsx)("div",{children:e.rowNum}),(0,g.jsx)("div",{className:e.pigeon_loft?"":"red-sign",children:e.pigeon_loft||"\u672a\u611f\u61c9\u638c\u4e2d\u5bf6\u6642\u9593"}),(0,g.jsx)("div",{children:e.pigeon_number}),(0,g.jsx)("div",{children:e.pigeon_aluminum}),(0,g.jsx)("div",{className:e.get_web_Induction_date?"":"red-sign",children:e.reset_get_web_Induction_date}),(0,g.jsx)("div",{className:e.reset_error_range_time.split("|").at(0),children:e.reset_error_range_time.split("|").at(1)}),(0,g.jsx)("div",{className:e.induction_date?"":"red-sign",children:e.reset_induction_date}),(0,g.jsx)("div",{children:e.aluminum_UID}),(0,g.jsxs)("div",{children:[e.latitude,"\xa0\xa0",e.longitude]})]},n)})):[]}},C),{},{postNext:function(e){return t(l.eQ,"setCurrentBackPaginationOptions",e)}});return clearInterval(f.current),f.current=setInterval((function(){0===v?(t(_.eQ,"setReflashCount",10),t(l.eQ,"getPigeonInductionData",{club:D,date:h||N})):t(_.eQ,"setReflashCount",v-1)}),1e3),(0,i.useEffect)((function(){var e=I.pages,a=I.partPage,i=I.pageSize,u=(0,c._T)(function(){var e=w?n.filter(b,(function(e){return e.pigeon_loft.match(w)})):b;return n.sort((0,r.Z)(e),(function(e,n){return{"":e.rowNum-n.rowNum,rowNum:Z?n.rowNum-e.rowNum:e.rowNum-n.rowNum,pigeon_number:Z?parseInt(e.pigeon_loft)-parseInt(n.pigeon_loft):parseInt(n.pigeon_loft)-parseInt(e.pigeon_loft),reset_get_web_Induction_date:Z?+new Date(e.reset_get_web_Induction_date)-+new Date(n.reset_get_web_Induction_date):+new Date(n.reset_get_web_Induction_date)-+new Date(e.reset_get_web_Induction_date),reset_error_range_time:function(){var t=e.reset_error_range_time.split("|").at(-1),r=n.reset_error_range_time.split("|").at(-1);return Z?parseInt(t)-parseInt(r):parseInt(r)-parseInt(t)}(),reset_induction_date:Z?+new Date("\u8acb\u611f\u61c9\u638c\u4e2d\u5bf6"===e.reset_induction_date?new Date:e.reset_induction_date)-+new Date("\u8acb\u611f\u61c9\u638c\u4e2d\u5bf6"===n.reset_induction_date?new Date:n.reset_induction_date):+new Date("\u8acb\u611f\u61c9\u638c\u4e2d\u5bf6"===n.reset_induction_date?new Date:n.reset_induction_date)-+new Date("\u8acb\u611f\u61c9\u638c\u4e2d\u5bf6"===e.reset_induction_date?new Date:e.reset_induction_date)}[O]}))}(),e,a,i),s=u.pageObj,m=u.renderItem;k((function(e){return(0,o.Z)((0,o.Z)({},e),{},{currentRenderItem:m})})),t(l.eQ,"setCurrentBackPaginationObj",(0,o.Z)({},s))}),[I,b,w,Z]),(0,i.useEffect)((function(){t(_.eQ,"setReflashCount",10)}),[w]),(0,i.useEffect)((function(){t(_.eQ,"setReflashCount",10),D&&N&&t(l.eQ,"getPigeonInductionData",{club:D,date:h||N})}),[h,Z]),(0,i.useEffect)((function(){if(D||N)return t(_.eQ,"setReflashCount",10),t(l.eQ,"setCurrentBackPaginationOptions",{pages:1,partPage:10,pageSize:10}),t(l.eQ,"getPigeonInductionData",{club:D,date:h||N}),function(){clearInterval(f.current)}}),[D,N]),(0,g.jsx)(m.Sn,(0,o.Z)({},B))}}}]);
//# sourceMappingURL=647.375ecd68.chunk.js.map