
function LogsGo() {

    var env = document.getElementById("EnvId").selectedIndex;
    if (env == 0 ) {
        alert('请选择环境！');
        return;
    }
	var domain = document.getElementById("domain").value;
    if (domain == 0 ) {
        alert('请选择应用！');
        return;
    }
	var machine = document.getElementById("machine").value;;
	if ( machine == 0) {
	    alert('请选择一台服务器！');
	    return;
	}
	var logsname = document.getElementById("LogsName").value;
	if (logsname == 0){
        alert('请选择日志！');
        return;
    }
    var loglevel = document.getElementById("loglevel").value;
	if ( loglevel == "") {
        alert('请至少选择一个log级别！');
        return;
    }
    var LogsTime = document.getElementById("d122").value;
    if ( LogsTime == "") {
        alert('请选择日期！');
        return;
    }
	var form1 = document.getElementById("form1");
	var viewtype = form1.view_type.value
	if (form1.view_type[0].checked ) {
        var line = document.getElementById("line").value;
        if (line == null || line == "") {
            alert('请填写行数！');
            return;
        }
	}else if (form1.view_type[1].checked ) {
        var min = document.getElementById("min").value;
        if (min == null || min == ""){
            alert('请填写分钟数！');
            return;
        }
	}else if (form1.view_type[2].checked ) {
	  var BeginTime = document.getElementById("BeginTime").value;
	  var EndTime = document.getElementById("EndTime").value;
	  if (BeginTime == null || BeginTime == "") {
          alert('请填写开始时间！');
          return;
      }else if (EndTime == null || EndTime == ""){
          alert('请填写结束时间！');
          return;
      }
    }
	//form1.submit();
    var data = {
	    'EnvId': env,
        'domain': domain,
        'machine': machine,
        'logsname':logsname,
        'LogsTime':LogsTime,
        'loglevel': loglevel,
        'viewtype':viewtype,
        'line':line,
        'min':min,
        'BeginTime':BeginTime,
        'EndTime':EndTime
    };
    $('#logs_content').html('')
	jQuery.ajax({
        url: 'getLogs',
        type: 'post',
        data: data,
        success: function(data){
        },
        error: function(error){
            console.log(error)
        }
    })
   }
   function changeRadio() {
        var form1 = document.getElementById("form1");
        var td1 = document.getElementById("td1");
        var td2 = document.getElementById("td2");
        var td3 = document.getElementById("td3");

        if (form1.view_type[0].checked) {
           td1.style.background="yellow";
           td1.style.fontWeight="bold";
           td2.style.background="white";
           td2.style.fontWeight="normal";
           td3.style.background="white";
           td3.style.fontWeight="normal";
        }

        if (form1.view_type[1].checked) {
           td1.style.background="white";
           td1.style.fontWeight="normal";
           td2.style.background="yellow";
           td2.style.fontWeight="bold";
           td3.style.background="white";
           td3.style.fontWeight="normal";
        }

        if (form1.view_type[2].checked) {
           td1.style.background="white";
           td1.style.fontWeight="normal";
           td2.style.background="white";
           td2.style.fontWeight="normal";
           td3.style.background="yellow";
           td3.style.fontWeight="bold";
        }
  }
  function focus_line() {
        var form1 = document.getElementById("form1");
        form1.view_type[0].checked = true;
        changeRadio();
  }
  function focus_min() {
        var form1 = document.getElementById("form1");
        form1.view_type[1].checked = true;
        changeRadio();
  }
  function getEnvList() {
        var EnvId = document.getElementById("EnvId").selectedIndex;
        if (EnvId == 0) {
           $("#domain option").remove();$("select[id='domain']").append("<option value=\"\">选择业务</option>");
           $("#machine option").remove();$("select[id='machine']").append("<option value=\"\">选择服务器</option>");
           $("#LogsName option").remove();$("select[id='LogsName']").append("<option value=\"\">选择日志</option>");
        }
        else {
           $.post("/getDomains",{ EnvId:$("select[id='EnvId'] option:selected").val() } ,function(data){
               $("#domain").html("");
               $("#domain").append("<option value=''>选择业务</option>");
                   if (typeof data !=='undefined' && data != null) {
                       $.each(data, function (idx, item) {
                           $("#domain").append("<option value=" + item.Domain + ">" + item.Domain + "</option>");
                       });
                       $("#domain").prop('selectedIndex', 0);
                   } else {
                       $("#EnvId").val('0');
                   }
           });
        }
  }
  function getMachineList() {
        var domain = document.getElementById("domain").selectedIndex;
        if (domain == 0) {
           $("#machine option").remove();$("select[id='machine']").append("<option value=\"\">选择服务器</option>");
        }
        else {
           $.post("getMachines",{
               EnvId:$("select[id='EnvId'] option:selected").val(),
               domain:$("select[id='domain'] option:selected").val()} ,function(data){
               $("#machine").html("");
               $("#machine").append("<option value=''>选择服务器</option>");
               if (typeof data !=='undefined' && data != null) {
                   $.each(data, function (idx, item) {
                       $("#machine").append("<option value=" + item.Machine + ">" + item.Machine + "</option>");
                   });
                   $("#machine").prop('selectedIndex', 0);
               } else {
                   $("#domain").val('0');
               }
           });
           }
  }
  function getLogsNameList() {
        var LogsName = document.getElementById("machine").selectedIndex;
        if (LogsName == 0) {
           $("#LogsName option").remove();$("select[id='LogsName']").append("<option value=\"\">选择日志</option>");
        }
        else {
           $.post("getLogsName",{
               EnvId:$("select[id='EnvId'] option:selected").val(),
               domain:$("select[id='domain'] option:selected").val(),
               machine:$("select[id='machine'] option:selected").val()} ,function(data){
               $("#LogsName").html("");
               $("#LogsName").append("<option value=''>选择服务器</option>");
               if (typeof data !=='undefined' && data != null) {
                   $.each(data, function (idx, item) {
                       $("#LogsName").append("<option value=" + item.Name + ">" + item.Name + "</option>");
                   });
                   $("#LogsName").prop('selectedIndex', 0);
               } else {
                   $("#machine").val('0');
               }
           });
           }
  }

