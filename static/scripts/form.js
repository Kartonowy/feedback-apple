b1.addEventListener("click", opinion_lesson);
b2.addEventListener("click", opinion_exams);
b3.addEventListener("click",  opinion_other);

let sliderValues = {};

         
function sliders(){
                
                        let c1 = document.createElement("div");
                        c1.style.width = "33%";
                        c1.id = "c1";


                        let c2 = document.createElement("div");
                        c2.style.width = "33%";
                        c2.id = "c2";

                        let c3 = document.createElement("div");
                        c3.style.width = "33%";
                        c3.id = "c3";

                        m.append(c1, c2, c3);
    
                        let slider = document.createElement("input");
                        slider.type = "range";
                        slider.id = "TeaSlider";
                        slider.min = 0;
                        slider.max = 10;
                        slider.step = 1;
                        c2.append(slider);
                
                        let output = document.createElement("p");
                        output.innerHTML = slider.value;
                        slider.oninput = function(){
                            output.innerHTML = this.value;
                        };
                        c2.append(output);
                
                
                };
//input sequence
//adding actions to history
function  opinion_lesson(){
                let selBut1 = this.textContent;
                sliderValues.first_choice = selBut1;
                console.log(selBut1);
                m.innerHTML = "";
                learning_pace = document.createElement("button");
                learning_pace.innerHTML = "Learning Pace";
                learning_pace.value = "Learning Pace";

                difficulty = document.createElement("button");
                difficulty.innerHTML = "Difficulty";
                difficulty.value = "Difficulty";

                comprehersability = document.createElement("button");
                comprehersability.innerHTML = "Comprehersability";
                comprehersability.value = "Comprehersability";

                learning_at = document.createElement("button");
                learning_at.innerHTML = "Learning Atmosphere";
                learning_at.value = "Learning Atmosphere";

                diversity = document.createElement("button");
                diversity.innerHTML = "Diversity";
                diversity.value = "Diversity";

                interactions = document.createElement("button");
                interactions.innerHTML = "Interactions";
                interactions.value = "Interactions";

                support = document.createElement("button");
                support.innerHTML = "Support";
                support.value = "Support";
                
                
                learning_pace.setAttribute("class", "chooseFeedback2");
                difficulty.setAttribute("class", "chooseFeedback2");
                comprehersability.setAttribute("class", "chooseFeedback2");
                learning_at.setAttribute("class", "chooseFeedback2");
                diversity.setAttribute("class", "chooseFeedback2");
                interactions.setAttribute("class", "chooseFeedback2");
                support.setAttribute("class", "chooseFeedback2");
                
                m.append(learning_pace);
                m.append(difficulty);
                m.append(comprehersability);
                m.append(learning_at);
                m.append(diversity);
                m.append(interactions);
                m.append(support);
                
                learning_pace.addEventListener("click", function ff1(){
                    let selBut2 = this.textContent;
                    sliderValues.second_choice = selBut2;
                    console.log(selBut2);
                    m.innerHTML = "";
                    sliders();
                    let lable = document.createElement("label");
                    lable.for = "slider";
                    lable.textContent = "Too Fast";
                        
                    let span = document.createElement("span");
                    span.textContent = "Too Slow";
                        
                     
                    c1.append(lable);
                    
                    c3.append(span);
                     
                });
                
                comprehersability.addEventListener("click", function ff3(){
                    let selBut2 = this.textContent;
                    sliderValues.second_choice = selBut2;
                    console.log(selBut2);
                    m.innerHTML = "";
                    let lable = document.createElement("label");
                    lable.for = "slider";
                    lable.textContent ="Not understandable" ;
                        
                    let span = document.createElement("span");
                    span.textContent = "Understandable";
                        
                    
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                     
                });
                
                difficulty.addEventListener("click", function ff2(){
                    let selBut2 = this.textContent;
                    sliderValues.second_choice = selBut2;
                    console.log(selBut2);
                    m.innerHTML = "";
                    let lable = document.createElement("label");
                    lable.for = "slider";
                    lable.textContent = "Undemanding";
                        
                    let span = document.createElement("span");
                    span.textContent = "Too Difficult";
                        
                        
                    sliders();    
                    c1.append(lable);
                    c3.append(span);
                   
                });
                
                
                
                learning_at.addEventListener("click", function ff4(){
                    let selBut2 = this.textContent;
                    sliderValues.second_choice = selBut2;
                    console.log(selBut2);
                    m.innerHTML = "";
                    let lable = document.createElement("label");
                    lable.for = "slider";
                    lable.textContent = "Too Quiet" ;
                        
                    let span = document.createElement("span");
                    span.textContent = "Too Loud";
                        
                        
                    sliders();    
                    c1.append(lable);
                    c3.append(span);
                    
                });
                
                diversity.addEventListener("click", function ff5(){
                    let selBut2 = this.textContent;
                    sliderValues.second_choice = selBut2;
                    console.log(selBut2);
                    m.innerHTML = "";
                    let lable = document.createElement("label");
                    lable.for = "slider";
                    lable.textContent = "Monotone";
                        
                    let span = document.createElement("span");
                    span.textContent ="Too Diverse" ;
                        
                        
                    sliders();    
                    c1.append(lable);
                    c3.append(span);
                    
                });
                
                
                interactions.addEventListener("click", function ff6(){
                    let selBut2 = this.textContent;
                    sliderValues.second_choice = selBut2;
                    console.log(selBut2);
                    m.innerHTML = "";
                    let a = document.createElement("button");
                    a.setAttribute("class", "chooseFeedback3");
                    a.innerHTML = "Frequency";
                    
                    
                    let b = document.createElement("button");
                    b.setAttribute("class", "chooseFeedback3");
                    b.innerHTML = "Way of Interaction";
                    
                    m.append(a);
                    m.append(b);
                    
                    a.addEventListener("click", function ffa(){
                        let selBut3 = this.textContent;
                        sliderValues.third_choice = selBut3;
                        console.log(selBut3);
                        m.innerHTML = "";
                        let lable = document.createElement("label");
                        lable.for = "slider";
                        lable.textContent = "Too Rarely" ;
                        let span = document.createElement("span");
                        span.textContent ="Too Often";
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                       
                    });
                    b.addEventListener("click", function ffb(){
                        let selBut3 = this.textContent;
                        sliderValues.third_choice = selBut3;
                        console.log(selBut3);
                        m.innerHTML = "";
                        let lable = document.createElement("label");
                        lable.for = "slider";
                        lable.textContent ="Unpolite" ;
                        let span = document.createElement("span");
                        span.textContent = "Polite";
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                         
                    });
                    
                });
                
                support.addEventListener("click", function ff7(){
                    let selBut2 = this.textContent;
                    sliderValues.second_choice = selBut2;
                    console.log(selBut2);
                    m.innerHTML = "";
                    let a = document.createElement("button");
                    a.setAttribute("class", "chooseFeedback3");
                    a.innerHTML = "Frequency";
                    
                    let b = document.createElement("button");
                    b.setAttribute("class", "chooseFeedback3");
                    b.innerHTML = "Quality of Support";
                    
                    m.append(a);
                    m.append(b);
                    
                    a.addEventListener("click", function ffa(){
                        let selBut3 = this.textContent;
                        sliderValues.third_choice = selBut3;
                        console.log(selBut3);
                        m.innerHTML = "";
                        let lable = document.createElement("label");
                        lable.for = "slider";
                        lable.textContent = "Too Rarely" ;
                        let span = document.createElement("span");
                        span.textContent = "Too Often";
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                         
                    });
                    b.addEventListener("click", function ffb(){
                        let selBut3 = this.textContent;
                        sliderValues.third_choice = selBut3;
                        console.log(selBut3);
                        m.innerHTML = "";
                        let lable = document.createElement("label");
                        lable.for = "slider";
                        lable.textContent = "Low" ;
                        let span = document.createElement("span");
                        span.textContent = "High";
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                        
                    });
                    
                });
            }
function  opinion_exams(){
                let selBut1 = this.textContent;
                sliderValues.first_choice = selBut1;
                console.log(selBut1);
                m.innerHTML = "";
                
                task_def = document.createElement("button");
                task_def.innerHTML = "Task Definition";
                task_def.value = "Task Definition";

                circumference = document.createElement("button");
                circumference.innerHTML = "Circumference";
                circumference.value = "Circumference";

                preparation = document.createElement("button");
                preparation.innerHTML = "Preparation";
                preparation.value = "Preparation";
                
                task_def.setAttribute("class", "chooseFeedback");
                circumference.setAttribute("class", "chooseFeedback");
                preparation.setAttribute("class", "chooseFeedback");
                m.append(task_def);
                m.append(circumference);
                m.append(preparation);
                
                
                
                    task_def.addEventListener("click", function ff1(){
                        let selBut2 = this.textContent;
                        sliderValues.second_choice = selBut2;
                        console.log(selBut2);

                        m.innerHTML = "";
                        let lable = document.createElement("label");
                        lable.for = "slider";
                        lable.textContent = "Not understandable" ;
                        
                        let span = document.createElement("span");
                        span.textContent =  "Understandable";
                        
                        
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                     
                        
                    });
                
                    circumference.addEventListener("click", function ff2(){
                        let selBut2 = this.textContent;
                        sliderValues.second_choice = selBut2;
                        console.log(selBut2);
                        m.innerHTML = "";
                        let lable = document.createElement("label");
                        lable.for = "slider";
                        lable.textContent = "Inadequate" ;
                        
                        let span = document.createElement("span");
                        span.textContent = "To wide";
                        
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                       
                        
                    });
                
                    preparation.addEventListener("click", function ff2(){
                        let selBut2 = this.textContent;
                        sliderValues.second_choice = selBut2;
                        console.log(selBut2);
                        m.innerHTML = "";
                        let lable = document.createElement("label");
                        lable.for = "slider";
                        lable.textContent =  "Poorly prepared";
                        
                        let span = document.createElement("span");
                        span.textContent ="Prepared";
                        
                        
                        sliders();    
                        c1.append(lable);
                        c3.append(span);
                         
                    });
               
            };
function opinion_other(){
    let selectedButton1Value = this.textContent;
    sliderValues.first_choice = selectedButton1Value;
    console.log(selectedButton1Value);
    m.innerHTML = "";
    let container = document.createElement("div");
    container.id = "con";
    m.append(container);
    let spillTea = document.createElement("h1");
    spillTea.innerHTML = "Don't spill the tea!";
    let decorum = document.createElement("h5");
    decorum.innerHTML = "We kindly ask to maintain decorum and respect. No unpolite questions or rumours are meant to be spread.";
    
    container.append(spillTea);
    container.append(decorum);
    
    let area = document.createElement("textarea"); 
    area.id = "teaArea"; 
    area.placeholder = "Write your feedback here..."; 
    
    container.append(area); 
    
}


document.addEventListener("input", function(event){
    
    if(event.target.type === 'range'){
        sliderValues[event.target.id] = event.target.value;
        console.log(event.target.value);
    }  
});

function getSliderValues() {
    
    sliderValues.first_choice.value;
    if(sliderValues.second_choice != undefined){
        sliderValues.second_choice.value;
    } else {
        sliderValues.second_choice == "";
    }

    if(sliderValues.third_choice != null){
        sliderValues.third_choice.value;
    } else {
        sliderValues.third_choice == "";
    }
    if (document.getElementById("TeaSlider") == undefined) {
        sliderValues.rating = document.querySelector("#teaArea").value;
    } else {
        sliderValues.rating = document.getElementById("TeaSlider").value;
    }

    return JSON.stringify(sliderValues);
};

sendButton.addEventListener("click", function(event){
    event.preventDefault();
    let sliderData = getSliderValues();
    console.log("slider values all hierarchy: ", sliderData); //slider values:  {"TeaSlider":"6"} 
    fetch("/rate/",
        {
            method: "POST",
            headers: {
                "Conntent-type": "application/json; charset=UTF-8"
            },
            body: sliderData,
            redirect: 'follow',
        }).then(response => {
            if (response.headers.get('HX-Redirect')) {
                window.location.href = response.headers.get('HX-Redirect');
            }
        })
});


       
