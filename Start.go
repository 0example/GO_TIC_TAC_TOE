package main

import (
	"fmt"
	"net/http"
)


func Home(w http.ResponseWriter, r *http.Request) {
	html := `<head>	
				
                    <script src='//ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js'></script>
              </head>    
                  <html>

					<body style=background-color:black>
					
<div id='result' style=background-color:black;border-color:black;color:white;position:absolute;font-size:2em;left:15em;top:0em;><h3></h3></div><br><br>


						

<form action="{{.html}}">

<button type="submit" id="re" style=position:absolute;left:58em;top:2em;width:11em;height:5em;color:white;background-color:black>Retry</button>

</form>            
<form action="receive" method="post" id="cf1">
<button type="button"  id="cl1" style=color:black;background-color:white;position:absolute;top:10em;left:37em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf2">
<button type="button" id="cl2" style=color:black;background-color:white;position:absolute;top:10em;left:48em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf3">
<button type="button" id="cl3" style=color:black;background-color:white;position:absolute;top:10em;left:59em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf6">
<button type="button" id="cl6" style=color:black;background-color:white;position:absolute;top:20em;left:59em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf5">
<button type="button" id="cl5" style=color:black;background-color:white;position:absolute;top:20em;left:48em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf4">
<button type="button" id="cl4" style=color:black;background-color:white;position:absolute;top:20em;left:37em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf7">
<button type="button" id="cl7" style=color:black;background-color:white;position:absolute;top:30em;left:37em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf8">
<button type="button" id="cl8" style=color:black;background-color:white;position:absolute;top:30em;left:48em;width:10em;height:9em;></button>
</form>

<form action="receive" method="post" id="cf9">
<button type="button" id="cl9" style=color:black;background-color:white;position:absolute;left:59em;width:10em;height:9em;top:30em;></button>
</form>

                  </body></html>

<script>



$('#cl1').click(function(){

		if($('#result h2').text()==""){

		if (document.getElementById("cl1").textContent==""){
			document.getElementById("cl1").innerHTML="X";

			if(document.getElementById("cl2").textContent=="X"){
				if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";
				}else if (document.getElementById("cl2").textContent==""){
         		 	document.getElementById("cl2").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}

			}else if(document.getElementById("cl4").textContent=="X"){
				if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
				}else if (document.getElementById("cl2").textContent==""){
         			 document.getElementById("cl2").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}

			}else if(document.getElementById("cl3").textContent=="X" && document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";
				}else if (document.getElementById("cl2").textContent==""){
         			 document.getElementById("cl2").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}

			}else if(document.getElementById("cl7").textContent=="X" && document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";
				}else if (document.getElementById("cl2").textContent==""){
         			 document.getElementById("cl2").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
	
			}else if(document.getElementById("cl9").textContent=="X" && document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}else if (document.getElementById("cl2").textContent==""){
         		 document.getElementById("cl2").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
	
			}else if (document.getElementById("cl2").textContent==""){
         		 document.getElementById("cl2").innerHTML="O";

        	}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl4").innerHTML="O";

			}else if(document.getElementById("cl5").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}
		}else{
			alert("Please try again");
		}
})



$('#cl2').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl2").textContent==""){
			document.getElementById("cl2").innerHTML="X";

			if(document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         		 	document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl5").textContent=="X" && document.getElementById("cl2").textContent=="X"){
				if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl8").textContent=="X" && document.getElementById("cl2").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        	}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl4").innerHTML="O";

			}else if(document.getElementById("cl5").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}
		}else{
			alert("Please try again")
		}
})



$('#cl3').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl3").textContent==""){

			document.getElementById("cl3").innerHTML="X";

			if(document.getElementById("cl2").textContent=="X"){
				if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
	
				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl6").textContent=="X"){
				if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         		 	document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl3").textContent=="X" && document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl9").textContent=="X" && document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}	
			}else if(document.getElementById("cl7").textContent=="X" && document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}	
			}else if(document.getElementById("cl5").textContent=="X" && document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}	
			}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        	}else if(document.getElementById("cl2").textContent==""){
				document.getElementById("cl2").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl4").innerHTML="O";

			}else if(document.getElementById("cl5").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}
		}else{
			alert("Please try again");
		}
})

$('#cl4').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl4").textContent==""){

			document.getElementById("cl4").innerHTML="X";

			if(document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";
	
				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl7").textContent=="X"){
				if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl5").textContent=="X" && document.getElementById("cl4").textContent=="X"){
				if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl6").textContent=="X" && document.getElementById("cl4").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        	}else if(document.getElementById("cl2").textContent==""){
				document.getElementById("cl2").innerHTML="O";

			}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl5").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}	
		}else{
			alert("Please try again");
		}
})

$('#cl5').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl5").textContent==""){

			document.getElementById("cl5").innerHTML="X";

			if(document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl2").textContent=="X"){
				if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}
			}else if(document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
				}
			}else if(document.getElementById("cl4").textContent=="X"){
				if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";
				}
			}else if(document.getElementById("cl6").textContent=="X"){
				if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";
				}
			}else if(document.getElementById("cl7").textContent=="X"){
				if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";
				}
			}else if(document.getElementById("cl8").textContent=="X"){
				if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";
				}
			}else if(document.getElementById("cl9").textContent=="X"){
				if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";
				}
			}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        	}else if(document.getElementById("cl2").textContent==""){
				document.getElementById("cl2").innerHTML="O";

			}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl4").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}
		}else{
			alert("Please try again");
		}
})

$('#cl6').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl6").textContent==""){

			document.getElementById("cl6").innerHTML="X";

			if(document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl9").textContent=="X"){
				if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl5").textContent=="X" && document.getElementById("cl6").textContent=="X"){
				if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";
	
				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
	
				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl4").textContent=="X" && document.getElementById("cl6").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        	}else if(document.getElementById("cl2").textContent==""){
				document.getElementById("cl2").innerHTML="O";

			}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl4").innerHTML="O";

			}else if(document.getElementById("cl5").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}
		}else{
			alert("Please try again");
		}
})

$('#cl7').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl7").textContent==""){

			document.getElementById("cl7").innerHTML="X";

			if(document.getElementById("cl4").textContent=="X"){
				if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl8").textContent=="X"){
				if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl7").textContent=="X" && document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         		 	document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl9").textContent=="X" && document.getElementById("cl7").textContent=="X"){
				if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}	
			}else if(document.getElementById("cl7").textContent=="X" && document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}	
			}else if(document.getElementById("cl7").textContent=="X" && document.getElementById("cl5").textContent=="X"){
				if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
         			 document.getElementById("cl1").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}	
			}else if (document.getElementById("cl1").textContent==""){
         		 document.getElementById("cl1").innerHTML="O";

        	}else if(document.getElementById("cl2").textContent==""){
				document.getElementById("cl2").innerHTML="O";

			}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl5").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}	
		}else{
			alert("Please try again");
		}
					
})

$('#cl8').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl8").textContent==""){

			document.getElementById("cl8").innerHTML="X";

			if(document.getElementById("cl7").textContent=="X"){
				if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
        			 document.getElementById("cl1").innerHTML="O";

       			}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl9").textContent=="X"){
				if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
        			 document.getElementById("cl1").innerHTML="O";

       			}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl8").textContent=="X" && document.getElementById("cl2").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
        			 document.getElementById("cl1").innerHTML="O";

       			}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
					document.getElementById("cl9").innerHTML="O";
				}
			}else if(document.getElementById("cl8").textContent=="X" && document.getElementById("cl5").textContent=="X"){
				if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";
				}else if (document.getElementById("cl1").textContent==""){
        			 document.getElementById("cl1").innerHTML="O";

       			}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
				}
			}else if (document.getElementById("cl1").textContent==""){
        		 document.getElementById("cl1").innerHTML="O";

       		}else if(document.getElementById("cl2").textContent==""){
				document.getElementById("cl2").innerHTML="O";

			}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl4").innerHTML="O";

			}else if(document.getElementById("cl5").textContent==""){
				document.getElementById("cl5").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl9").textContent==""){
				document.getElementById("cl9").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}
		}else{
			alert("Please try again");
		}
})

$('#cl9').click(function(){

if($('#result h2').text()==""){

if (document.getElementById("cl9").textContent==""){
			document.getElementById("cl9").innerHTML="X";

			if(document.getElementById("cl6").textContent=="X"){
				if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";
				}else if (document.getElementById("cl5").textContent==""){
         			 document.getElementById("cl5").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}
			}else if(document.getElementById("cl8").textContent=="X"){
				if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
				}else if (document.getElementById("cl5").textContent==""){
         			 document.getElementById("cl5").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";
	
				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}
			}else if(document.getElementById("cl9").textContent=="X" && document.getElementById("cl3").textContent=="X"){
				if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";
				}else if (document.getElementById("cl5").textContent==""){
         			 document.getElementById("cl5").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}
			}else if(document.getElementById("cl9").textContent=="X" && document.getElementById("cl7").textContent=="X"){
				if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}else if (document.getElementById("cl5").textContent==""){
         			 document.getElementById("cl5").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}	
			}else if(document.getElementById("cl9").textContent=="X" && document.getElementById("cl1").textContent=="X"){
				if(document.getElementById("cl5").textContent==""){
					document.getElementById("cl5").innerHTML="O";
				}else if (document.getElementById("cl5").textContent==""){
         		 document.getElementById("cl5").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";
	
				}else if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";
	
				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";
	
				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
				}	
			}else if(document.getElementById("cl9").textContent=="X" && document.getElementById("cl5").textContent=="X"){
				if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";
				}else if (document.getElementById("cl5").textContent==""){
         			 document.getElementById("cl5").innerHTML="O";

        		}else if(document.getElementById("cl2").textContent==""){
					document.getElementById("cl2").innerHTML="O";

				}else if(document.getElementById("cl3").textContent==""){
					document.getElementById("cl3").innerHTML="O";

				}else if(document.getElementById("cl4").textContent==""){
					document.getElementById("cl4").innerHTML="O";

				}else if(document.getElementById("cl1").textContent==""){
					document.getElementById("cl1").innerHTML="O";

				}else if(document.getElementById("cl6").textContent==""){
					document.getElementById("cl6").innerHTML="O";

				}else if(document.getElementById("cl7").textContent==""){
					document.getElementById("cl7").innerHTML="O";

				}else if(document.getElementById("cl8").textContent==""){
					document.getElementById("cl8").innerHTML="O";
			}	
			}else if (document.getElementById("cl5").textContent==""){
         		 document.getElementById("cl5").innerHTML="O";

        	}else if(document.getElementById("cl2").textContent==""){
				document.getElementById("cl2").innerHTML="O";

			}else if(document.getElementById("cl3").textContent==""){
				document.getElementById("cl3").innerHTML="O";

			}else if(document.getElementById("cl4").textContent==""){
				document.getElementById("cl4").innerHTML="O";

			}else if(document.getElementById("cl1").textContent==""){
				document.getElementById("cl1").innerHTML="O";

			}else if(document.getElementById("cl6").textContent==""){
				document.getElementById("cl6").innerHTML="O";

			}else if(document.getElementById("cl7").textContent==""){
				document.getElementById("cl7").innerHTML="O";

			}else if(document.getElementById("cl8").textContent==""){
				document.getElementById("cl8").innerHTML="O";
			}

		} else {

			alert("try another squre");
		}	
		}else{
			alert("Please try again");
		}
					
})




                         $('#cl1').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
                                  $('#result').html(data);
                               },
                             });
                          });


                         $('#cl2').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
                                	$('#result').html(data); 
								},
                             });
                          });


							$('#cl3').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
                                	$('#result').html(data); 
								},
                             });
                          });

							$('#cl4').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
									$('#result').html(data);
								},
                             });
                          });

							$('#cl5').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
									$('#result').html(data);
                               },
                             });
                          });

							$('#cl6').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
                                  $('#result').html(data);
                               },
                             });
                          });

							$('#cl7').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
									$('#result').html(data);
                               },
                             });
                          });

							$('#cl8').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
									$('#result').html(data);
                               },
                             });
                          });

							$('#cl9').click(function () {
							var s= $('#cl1').text();
							var s2= $('#cl2').text();
							var s3= $('#cl3').text();
							var s4= $('#cl4').text();
							var s5= $('#cl5').text();
							var s6= $('#cl6').text();
							var s7= $('#cl7').text();
							var s8= $('#cl8').text();
							var s9= $('#cl9').text();
                             $.ajax({
                               url: 'receive',
                               type: 'post',
                               dataType: 'html',
                               data : { cf1:s,cf2:s2,cf3:s3,cf4:s4,cf5:s5,cf6:s6,cf7:s7,cf8:s8,cf9:s9},
                               success : function(data) {
									$('#result').html(data);
                               },
                             });
                          });



</script>`

	w.Write([]byte(fmt.Sprintf(html)))

}

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		data1 := r.FormValue("cf1")
		data2 := r.FormValue("cf2")
		data3 := r.FormValue("cf3")

		data4 := r.FormValue("cf4")
		data5 := r.FormValue("cf5")
		data6 := r.FormValue("cf6")
		//
		data7 := r.FormValue("cf7")
		data8 := r.FormValue("cf8")
		data9 := r.FormValue("cf9")

		if(data1==data2 && data3==data1 ){
			if(data1!=""){
				if (data1=="X") {
					fmt.Println("You Win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lost<h2>"))
				}
				fmt.Println("1")
			}



		}else if(data4==data5 && data4==data6 ){
			if(data4!=""){
				if(data4=="X"){
					fmt.Println("You Win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lose<h2>"))
				}
				fmt.Println("2")
			}

		}else if (data7==data8 && data8==data9 ) {

			if(data7!=""){
				if(data7=="X"){
					fmt.Println("You Win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lose<h2>"))
				}
				fmt.Println("3")
			}

		}else if(data1==data4 && data4==data7 ){
			if(data4!=""){
				if (data4=="X") {
					fmt.Println("You win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lose<h2>"))
				}
				fmt.Println("4")
			}

		}else if(data2==data5 && data2==data8 ){
			if(data2!=""){
				if(data2=="X"){
					fmt.Println("You Win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lose<h2>"))
				}
				fmt.Println("5")
			}

		}else if(data3==data6 && data6==data9  ){
			if(data3!=""){
				if(data3=="X"){
					fmt.Println("You Win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lose<h2>"))
				}
				fmt.Println("6")
			}

		}else if(data7==data5 && data7==data3 ){
			if(data7!=""){
				if(data7=="X"){
					fmt.Println("You Win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lose<h2>"))
				}
				fmt.Println("7")
			}

		}else if(data1==data5 && data9==data1 ){
			if(data5!=""){
				if (data5=="X") {
					fmt.Println("You Win")
					w.Write([]byte("<h2>You Win<h2>"))
				}else {
					fmt.Println("You Lose")
					w.Write([]byte("<h2>You Lose<h2>"))
				}
				fmt.Println("8")

			}

		}else if (data1!="" && data2!="" && data3!="" && data4!="" && data5!="" && data6!="" && data7!="" && data8!="" && data9!="") {
			fmt.Println("Draw")
			w.Write([]byte("<h2>Draw<h2>"))
		}


	}
}

func main() {
	// http.Handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/receive", receiveAjax)

	http.ListenAndServe(":8080", mux)
}