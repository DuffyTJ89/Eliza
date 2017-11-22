const conversation = $("#list"); // $ means jquery, it won't get the item without the #
const userInput = $("#userInput");
 console.log(conversation);


userInput.keypress(function(event){
    console.log("working");
    if(event.keyCode !== 13){ // 13 is the keycode for Enter
        return; // do nothing unless the key is enter
    }
    event.preventDefault(); // prevents the form default behaviour which would refresh the page.
    const text = userInput.val();

    console.log(text);
    userInput.val(""); // set it to nothing, .val() is like a getter, .val(" ") is like a setter

    // trim removes all spaces from either side,
    // if there's no text left, the user doesn't have a question.
    if(text.trim() == ""){
    //if(!text.trim()){ 
        return;
    }

    // a query parameter user-input is expected
    queryParameters = {
        "userInput" : text
    }
    //<li class="list-group-item list-group-item-success">First item</li>
    //<li class="list-group-item list-group-item-success text-right">
    //conversation.append("<li id='user' class='list-group'>" + text + "<li class='list-group'>");
    conversation.append('<li id="user" class="list-group-item list-group-item-success text-right">' + text + "  : User" + '<li class="list-group">');


    $.get("/chat", queryParameters).done(function(resp){
        // this code will execute when the request gets a response.
        setTimeout(function(){ // wait 1 second then add element.
            conversation.append('<li id="eliza" class="list-group-item list-group-item-primary">'+"ELIZA :  "+  resp +  '<li class="list-group">');
        }, 1000);
        
    }).fail(function(){ // this will run whenever anything goes wrong.
        conversation.append("<li class='list-group'>Error :( </li class='list-group\'>");
    });

    window.scrollTo(0,document.body.scrollHeight); //scroll to the bottom so the latest chat is in view

    
});
