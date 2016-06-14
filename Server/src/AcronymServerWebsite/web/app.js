function Update(acronyms) {

    var length = acronyms.length
    var perGroup = 3;
    var groupLength = Math.ceil(acronyms.length / perGroup)
    $('#divAcronyms').empty()

    var child = ""
    for(var j=0;j<groupLength;j++) {
        var isRow = j % 2 == 0

        if(isRow) {
            child += '<div class="row">'
        }

        child += '<div class="col-lg-6">'

        for(var i=0;i<3;i++) {
            var index = j*perGroup+i 
            if(index >= length) {
                break;
            }

            var acronym = acronyms[index]

            var definition = acronym.Definition
            definition = definition.replace(/\n/g, "<br/>")

            child += "<h4>" + acronym.Acronym + "</h4>" +
                        "<h6>Language: " + acronym.Language + "</h6>" +
                        "<p>" + definition + "</p>" +
                        '<a href="'+ acronym.Url +'"> Link </a>';

        }

        child +='</div>'
        if(!isRow) {
            child += '</div>'
        }
    }

    $('#divAcronyms').append(child)
}

function Init() {

    $.ajax({
        url: "api/v1/Random/6",
        context: document.body
    }).success(function(data) {
        
        json = JSON.parse(data)
        Update(json)
    
    });
}

function SearchClick() {

    $.ajax({
        url: "api/v1/Acronym/" + $('#txtAcronym').val(),
        context: document.body
    }).success(function(data) {
        
        json = JSON.parse(data)
        Update(json)
    
    });
}

$(function() {

    Init();
    $('#btnSearch').click(SearchClick)

});