const api_url = 
  "http://localhost:1234/api/alllocations";

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
    <tr>
    <th>Location-Name</th>
    <th>Location-Langitude</th>
    <th>Location-Latitude</th>
    <th>Map</th>
    <th>Visited</th>
    </tr>`;

    for (let r of data) {
      var checkbox = ""
      if (r.Visited) {
        checkbox = "checked"
      }
      tab += `<tr>
      <td>${r.Placename}</td>
      <td>${r.Lat}</td>
      <td>${r.Long}</td>
      <td id="maps${r.id}" style="height: 180px; width: 180px"></td>
      <td><input type="checkbox" ${checkbox} disabled style="hight: 60px; width: 60px"></td>
  
      </tr>`;
    }
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
