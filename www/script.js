function getIcpInfo() {
    let url = document.getElementById("url").value;
    let xhr = new XMLHttpRequest();
    
    // 设置 withCredentials 为 true
    xhr.withCredentials = true;
    
    xhr.onreadystatechange = function() {
      if (xhr.readyState === XMLHttpRequest.DONE) {
        if (xhr.status === 200) {
          let response = JSON.parse(xhr.responseText);
          let icpSerial = response.ICPSerial;
          let organization = response.Organization;
          let detectTime = new Date(response.detect_time * 1000).toLocaleString(); // 转换为本地时间
  
          if (response.isDomainICPOk === "True") {
            document.getElementById("result").value = `ICP备案号为：\n${icpSerial}\n备案主体为：\n${organization}\n查询时间为：\n${detectTime}`;
          } else {
            document.getElementById("result").value = "该网站没有ICP备案信息。";
          }
        } else {
          console.log("请求失败！");
        }
      } 
    };
    
    xhr.open("GET", "/geticpinfo?url=" + url);
    xhr.send();
  }
  