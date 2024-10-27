
const search = async (doc,padron_type)=>{
    const url="http://127.0.0.1:3001/api/"+padron_type+"/"+doc;
    console.log(url);
    const response=await fetch(url)
    .then(response => response.json())
    .then(data => data);

    // console.log(response);
    parseResponse(response);
};

const parseResponse=(data)=>{

    const responseEl=document.querySelector("#response");
    if (data.Doc!=""){
        content=`
        <div class="info">
            <h3>Datos Personales:</h3>
            <p><b>Doc:</b>${data.Doc}</p>
            <p><b>Apellido y nombre:</b>${data.Fullname}</p>
            <p><b>Domicilio:</b>${data.Address}</p>
            <p><b>Género:</b>${data.Gender}</p>
            <p><b>Distrito:</b>${data.DistritoName}</p>
            <h3>Lugar de votación</h3>
            <p><b>Establecimiento:</b>${data.EstablecimientoName}</p>
            <p><b>Dirección:</b>${data.EstablecimientoLocationAddress}</p>
            <p><b>Localidad:</b>${data.EstablecimientoLocationName}</p>
            <p><b>Mesa:</b>${data.Mesa}</p>
            <p><b>Órden:</b>${data.Orden}</p>
            <p><b>Cod. Circ.:</b>${data.CodCirc}</p>
        </div>
        `;
        responseEl.innerHTML=content;
    }else{
        responseEl.innerHTML="<h3 class='error'>La persona no se encuentra en el padrón.</h3>";
    }
};

document.addEventListener('DOMContentLoaded',async ()=>{
    
    const btnSearchEl = document.querySelector("#btn_search");
    const docInputEl = document.querySelector("#doc_input");
    const padronTypeSelectorEl = document.querySelector("#padron_type_selector");

    

    btnSearchEl.addEventListener('click',async (e)=>{
        e.preventDefault();
        await search(docInputEl.value,padronTypeSelectorEl.value);
    });

});