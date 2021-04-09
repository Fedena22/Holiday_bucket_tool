const api_url = 
  "http://localhost:3000/admin/alllocations";

  async function getapi(url){
    
    const respone = await fetch(url, {
      headers: {
        Authorization: 'Basic '+btoa('admin:admin'), 
      }
    })
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
    <th>Checked</th>
    </tr>`;

    for (let r of data) {
    // var map = new L.map('maps', {
    //     center: [r.Lat, r.Long],
    //     zoom: 18
    // });
    // const latitude = r.Lat;
    // const longitude = r.Long;
    // const titleServer1 = 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png';
    // const titleServer2 = 'https://cdn.lima-labs.com/{z}/{x}/{y}.png?free ';
    //   L.tileLayer(titleServer1, {
    //     attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    // }).addTo(map);
    // L.marker([latitude,longitude ]).addTo(map)
    // console.log(map);
      tab += `<tr>
      <td>${r.Placename}</td>
      <td>${r.Lat}</td>
      <td>${r.Long}</td>
      <td>map</td>
      <td>${r.Visited}</td>

      </tr>`;
    }
    document.getElementById("location").innerHTML = tab;
  }
