const conversation = $("#list"); // $ means jquery, it won't get the item without the #
const userInput = $("#userInput");
 console.log(conversation); //keep a record of the user side of the conversation in the console


userInput.keypress(function(event){
    console.log("working"); //log to the console that the enter key is working
    if(event.keyCode !== 13){ // 13 is the keycode for Enter
        return; // do nothing unless the key is enter
    }
    event.preventDefault(); // prevents the form default behaviour which would refresh the page.
    const text = userInput.val(); //set user input to text

    console.log(text); //user input to console for logging
    userInput.val(""); // set it to nothing, .val() is like a getter, .val(" ") is like a setter

    // trim removes all spaces from either side,
    // if there's no text left, the user doesn't have a question.
    if(text.trim() == ""){
        return;
    }

    // a query parameter user-input is expected
    queryParameters = {
        "userInput" : text
    }
    //add the next part of the conversation to the list group so it can be displayed
    conversation.append('<li id="user" class="list-group-item list-group-item-success text-right">' + text + "  : User" + '<li class="list-group">');


    $.get("/chat", queryParameters).done(function(resp){
        // this code will execute when the request gets a response.
        setTimeout(function(){ // wait 1 second then add Eliza's response.
            conversation.append('<li id="eliza" class="list-group-item list-group-item-primary">'+"ELIZA :  "+  resp +  '<li class="list-group">');
        }, 1000);
        
    }).fail(function(){ // this will run whenever anything goes wrong.
        conversation.append("<li class='list-group'>Error :( </li class='list-group\'>");
    });

    window.scrollTo(0,document.body.scrollHeight); //scroll to the bottom so the latest chat is in view

    
});
