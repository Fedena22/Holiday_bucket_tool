const api_url = 
  "http://127.0.0.1:1234/api/alllocations";

  async function getapi(url){
    
    const respone = await fetch(url)
    var data = await respone.json();
    console.log(data);
    if (respone) {
      hideloader();
    }
    show(data);
  }

  getapi(api_url);

  function hideloader() {
    document.getElementById('loading').style.display = 'none';
  }

  function show(data){
    let tab =
    `
    <thead>
    <tr>
    <th>Location-Name</th>
    <th>Map</th>
    <th>Visited</th>
    </tr>
    </thead>`;

    for (let r of data) {
      var checkbox = ""
      if (r.Visited) {
        checkbox = "checked"
      }
      tab += `<tbody><tr>
      <td style="width: 40%">${r.Placename}</td>
      <td id="maps${r.id}" style="height: 180px; width: 50%"></td>
      <td><input type="checkbox" ${checkbox}  style="hight: 180px; width: 10%"></td>
  
      </tr>
      </tbody>`;
    }
    tab += 
    `
    <tfoot>
    <tr>
    <th>Location-Name</th>
    <th>Map</th>
    <th>Visited</th>
    </tr>
    </tfoot>`;
    document.getElementById("location").innerHTML = tab;

  for (let r of data) {
    var maps = "maps" + r.id
    let map = new L.map(maps, {
      center: [r.Lat, r.Long],
      zoom: 18
  });
  const latitude = r.Lat;
  const longitude = r.Long;
  const titleServer1 = 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
  // const titleServer2 = 'https://cdn.lima-labs.com/{z}/{x}/{y}.png?free ';
    L.tileLayer(titleServer1, {
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
  }).addTo(map);
  L.marker([latitude,longitude ]).addTo(map)
    }

  }
  function addRow(tableID) {

    var table = document.getElementById(tableID).getElementsByTagName('tbody')[0];

    // var rowCount = table.rows.length;
    var row = table.insertRow(table.rows.length);

    var cell1 = row.insertCell(0);
    var element1 = document.createElement("input");
    element1.type = "text";
    element1.name="txtbox[]";
    element1.placeholder="Location name";
    cell1.appendChild(element1);

    var cell2 = row.insertCell(1);
    var element2 = document.createElement("input");
    element2.type ="text";
    element2.name ="latlong[]";
    element2.placeholder="Latitude / Longitude of the location"
    cell2.appendChild(element2);
}

function deleteRow(tableID) {
    try {
    var table = document.getElementById(tableID);
    var rowCount = table.rows.length;
    console.log(rowCount)
    for(var i=0; i<rowCount; i++) {
        var row = table.rows[i];
        var chkbox = row.cells[0].childNodes[0];
        if(null != chkbox && true == chkbox.checked) {
            table.deleteRow(i);
            rowCount--;
            i--;
        }


    }
    }catch(e) {
        alert(e);
    }
}

const newinput = document.querySelector("#newlocation");
if (newinput) {
  newinput.addEventListener("submit",function(e) {
    submitNewEntry(e, this);
  });
}

async function submitNewEntry(e, form) {
  e.preventDefault();
  const btnSubmit = document.getElementById("btnSubmit");
  btnSubmit.disable = true;
  setTimeout(() => btnSubmit.disable = false , 2000);

  const jsonFormData = buildJsonFormData(form);
  const headers = buildHeaders();
  // 2.4 Request & Response
  const response = await fetchService.performPostHttpRequest(`"http://127.0.0.1:1234/api/addlocations`, headers, jsonFormData); // Uses JSON Placeholder
  console.log(response);
  // 2.5 Inform user of result
  if(response)
      window.location = `/success.html?FirstName=${response.FirstName}&LastName=${response.LastName}&Email=${response.Email}&id=${response.id}`;
  else
      alert(`An error occured.`);
}

function buildHeaders() {
  const headers = {
      "Content-Type": "application/json",
  };
  return headers;
}

function buildJsonFormData(form) {
  const jsonFormData = { };
  for(const pair of new FormData(form)) {
      jsonFormData[pair[0]] = pair[1];
  }
  return jsonFormData;
}
/*--/Functions--*/