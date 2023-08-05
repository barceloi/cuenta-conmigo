package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	letterA = `
Auto / Car
                             _.-="_-         _
                         _.-="   _-          | ||"""""""---._______     __..
             ___.===""""-.______-,,,,,,,,,,,,\-''----" """""       """""  __'
      __.--""     __        ,'                   o \           __        [__|
 __-""=======.--""  ""--.=================================.--""  ""--.=======:
]       [w] : /        \ : |========================|    : /        \ :  [w] :
V___________:|          |: |========================|    :|          |:   _-"
 V__________: \        / :_|=======================/_____: \        / :__-"
 -----------'  "-____-"  \-------------------------------'  "-____-"

`
	letterB = `
 	Braquiosaurio

	  _
 .~q \,
{__,  \
    \' \
     \  \
      \  \
       \  \._            __.__
        \    ~-._  _.==~~     ~~--.._
         \        '                  ~-.
          \      _-   -_                \.
           \    /       }        .-    .  \
            \. |      /  }      (       ;  \
              \|     /  /       (       :    \
               \    |  /        |      /       \    =
                |     /\-.______.\     |^-.      \
                |   |/           (     |    .      \_
                |   ||            ~\   \      '._    \-.._____..----.._=__
                |   |/             _\   \      =~-.__________.-~~~~~~~~~'''
              .o'___/            .o______}`
	letterBOption2 = `
  Barco / Ship
  ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⠀⠤⠴⠶⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣾⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠂⠉⡇⠀⠀⠀⢰⣿⣿⣿⣿⣧⠀⠀⢀⣄⣀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢠⣶⣶⣷⠀⠀⠀⠸⠟⠁⠀⡇⠀⠀⠀⠀⠀⢹⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠘⠟⢹⣋⣀⡀⢀⣤⣶⣿⣿⣿⣿⣿⡿⠛⣠⣼⣿⡟⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣴⣾⣿⣿⣿⣿⢁⣾⣿⣿⣿⣿⣿⣿⡿⢁⣾⣿⣿⣿⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠸⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⡇⢸⣿⣿⣿⠿⠇⠀⠀⠀⠀
⠀⠀⠀⠳⣤⣙⠟⠛⢻⠿⣿⠸⣿⣿⣿⣿⣿⣿⣿⣇⠘⠉⠀⢸⠀⢀⣠⠀⠀⠀
⠀⠀⠀⠀⠈⠻⣷⣦⣼⠀⠀⠀⢻⣿⣿⠿⢿⡿⠿⣿⡄⠀⠀⣼⣷⣿⣿⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠈⣿⣿⣿⣶⣄⡈⠉⠀⠀⢸⡇⠀⠀⠉⠂⠀⣿⣿⣿⣧⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣿⣷⣤⣀⣸⣧⣠⣤⣴⣶⣾⣿⣿⣿⡿⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠟⠛⠉⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
  `
	letterBOption3 = `
  Ballena / Whale
                __   __
              __ \ / __
             /  \ | /  \
                 \|/
            _,.---v---._
   /\__/\  /            \
   \_  _/ /              \
     \ \_|           @ __|
      \                \_
       \     ,__/       /
     ~~~\~~~~~~~~~~~~~~/~~~~
  `
	letterC = `
  Camión / Truck
                         _____________________________________________________
                      |                                                     |
             _______  |                                                     |
            / _____ | |                                                     |
           / /(__) || |                                                     |
  ________/ / |OO| || |                                                     |
 |         |-------|| |                                                     |
(|         |     -.|| |_______________________                              |
 |  ____   \       ||_________||____________  |             ____      ____  |
/| / __ \   |______||     / __ \   / __ \   | |            / __ \    / __ \ |\
\|| /  \ |_______________| /  \ |_| /  \ |__| |___________| /  \ |__| /  \|_|/
   | () |                 | () |   | () |                  | () |    | () |
    \__/                   \__/     \__/                    \__/      \__/`

	letterD = `
Delfín / Dolphin

                                      .--.
                _______             .-"  .'
        .---u"""       """"---._  ."    %
      .'                        "--.    %
 __.--'  o                          "".. "
(____.                                  ":
 \----.__                                 ".
         \----------__                     ".
               ".   . ""--.                 ".
                 ". ".     ""-.              ".
                   "-.)        ""-.           ".
                                   "".         ".
                                      "".       ".
                                         "".      ".
                                            "".    ".
                      ^~^~^~^~^~^~^~^~^~^~^~^~^"".  "^~^~^~^~^
                                            ^~^~^~^  ~^~
                                                 ^~^~^~
`
	letterE = `
Elefante \ Elephant
                            _
                          .' \'.__
                         /      \ \'"-,
        .-''''--...__..-/ .     |      \
      .'               ; :'     '.  a   |
     /                 | :.       \     =\
    ;                   \':.      /  ,-.__;.-;'
   /|     .              '--._   /-.7\._..-;\
  ; |       '                |\-'      \  =|
  |/\        .   -' /     /  ;         |  =/
  (( ;.       ,_  .:|     | /     /\   | =|
   ) /  \     | \""\;     / |    | /   / =/
     | ::|    |      \    \ \    \ \--' =/
    /  '/\    /       )    |/     \-...-'
   /    | |   \    /-'    /;
   \  ,,/ |    \   D    .'  \
    '""'   \  nnh  D_.-'L__nnh
            '"""'

`
	letterF = `
  Flores / Flowers

                      _
                  _(_)_                          wWWWw   _
      @@@@       (_)@(_)   vVVVv     _     @@@@  (___) _(_)_
     @@()@@ wWWWw  (_)\    (___)   _(_)_  @@()@@   Y  (_)@(_)
      @@@@  (___)     \|/    Y    (_)@(_)  @@@@   \|/   (_)\
       /      Y       \|    \|/    /(_)    \|      |/      |
    \ |     \ |/       | / \ | /  \|/       |/    \|      \|/
jgs \\|//   \\|///  \\\|//\\\|/// \|///  \\\|//  \\|//  \\\|// 
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^`

	letterG = `
  Gato / Cat

      |\      _,,,---,,_
ZZZzz /,'.-'''    -.  ;-;;,_
     |,4-  ) )-,_. ,\ (  ''-'
    '---''(_/--'  '-'\_)  Felix Lee`
	letterH = `
  Helicóptero / Helicopter

     -----|-----
  *>=====[_]L)
        -'-\-
  `
	letterI = `
    Isla \ Island

                    
                           ,--.\'-. __
                        _,.'. \:/,"  '-._
                     ,-*" _,.-;-*'-.+"*._ )
                    ( ,."* ,-" / ''.  \.  '.
                   ,"   ,;"  ,"\../\  \:   \
                  (   ,"/   / \.,' :   ))  /
                   \  |/   / \.,'  /  // ,'
                    \_)\ ,' \.,'  (  / )/
                        '  \._,'   '"
                           \../
                           \../
                 ~        ~\../           ~~
          ~~          ~~   \../   ~~   ~      ~~
     ~~    ~   ~~  __...---\../-...__ ~~~     ~~
       ~~~~  ~_,--'        \../      '--.__ ~~    ~~
   ~~~  __,--'              '"             '--.__   ~~~
~~  ,--'                                         '--.
   '------......______             ______......------' ~~
 ~~~   ~    ~~      ~ '''''---"""""  ~~   ~     ~~
        ~~~~    ~~  ~~~~       ~~~~~~  ~ ~~   ~~ ~~~  ~
     ~~   ~   ~~~     ~~~ ~         ~~       ~~   SSt
              ~        ~~       ~~~       ~`
	letterJ = `
  Joya / Jewel
                              .     '     ,
                                _________
                            _ /_|_____|_\ _
                              '. \   / .'
                                '.\ /.'
                                  '.'

  `
	letterK = `
  Koala
               )    (   |
               )    (  /    .-
        _ ,---. _   ( /    /
      (~-| . . |-~)  V    /
       \._  0  _,/       /
        / \-^-'\-._     /
       '           \-. (
      :               )E
      :          ,---' (
       .            )E (
        '._____,---'   (
               )       (
               )       (
               )       (
               )       (


  `
	letterL = `
  Lupa / Magnifying glass

            ______              
         .-'        -.           
       .'            \.         
      /                \        
     ;                  ;       
     |                  |;       
     ;                 ;|
     '\               / ;       
      \\.           .' /        
       \.\-._____.-' .'         
         / /\_____.-'           
        / / /                   
       / / /
      / / /
     / / /
    / / /
   / / /
  / / /
 / / /
/ / /
\/_/`

	letterM = `
Mapa / Map
 ____________________________________________________________________
 / \-----     ---------  -----------     -------------- ------    ----\
 \_/__________________________________________________________________/
 |~ ~~ ~~~ ~ ~ ~~~ ~ _____.----------._ ~~~  ~~~~ ~~   ~~  ~~~~~ ~~~~|
 |  _   ~~ ~~ __,---'_       "         \. ~~~ _,--.  ~~~~ __,---.  ~~|
 | | \___ ~~ /      ( )   "          "   \-.,' (') \~~ ~ (  / _\ \~~ |
 |  \    \__/_   __(( _)_      (    "   "     (_\_) \___~ \-.___,'  ~|
 |~~ \     (  )_(__)_|( ))  "   ))          "   |    "  \ ~~ ~~~ _ ~~|
 |  ~ \__ (( _( (  ))  ) _)    ((     \\//    " |   "    \_____,' | ~|
 |~~ ~   \  ( ))(_)(_)_)|  "    ))    //\\ " __,---._  "  "   "  /~~~|
 |    ~~~ |(_ _)| | |   |   "  (   "      ,-'~~~ ~~~ \-.   ___  /~ ~ |
 | ~~     |  |  |   |   _,--- ,--. _  "  (~~  ~~~~  ~~~ ) /___\ \~~ ~|
 |  ~ ~~ /   |      _,----._,'\--'\.\-._  \._~~_~__~_,-'  |H__|  \ ~~|
 |~~    / "     _,-' / '\ ,' / _'  \'.---.._          __        " \~ |
 | ~~~ / /   .-' , / ' _,'_  -  _ '- _'._ ''.'-._    _/- '--.   " " \~|
 |  ~ / / _-- '---,~.-' __   --  _,---.  '-._   _,-'- / ' \ \_   " |~|
 | ~ | | -- _    /~/  '-_- _  _,' '  \ \_'-._,-'  / --   \  - \_   / |
 |~~ | \ -      /~~| "     ,-'_ /-  '_ ._'._'-...._____...._,--'  /~~|
 | ~~\  \_ /   /~~/    ___  '---  ---  - - ' ,--.     ___        |~ ~|
 |~   \      ,'~~|  " (o o)   "         " " |~~~ \_,-' ~ '.     ,'~~ |
 | ~~ ~|__,-'~~~~~\    \"/      "  "   "    /~ ~~   O ~ ~~'-.__/~ ~~~|
 |~~~ ~~~  ~~~~~~~~'.______________________/ ~~~    |   ~~~ ~~ ~ ~~~~|
 |____~jrei~__~_______~~_~____~~_____~~___~_~~___~\_|_/ ~_____~___~__|
 / \----- ----- ------------  ------- ----- -------  --------  -------\
 \_/__________________________________________________________________/

`
	letterN = `
Nubes / Clouds
              (  ).                   _           
             (     ).              .:('  )'.       
)           _(       ''.          :(   .    )      
        .=('(      .   )     .--  '.  (    ) )      
       ((    (..__.:'-'   .+(   )   '' _'  ) )                 
 .      (       ) )       (   .  )     (   )  ._   
  )      ' __.:'   )     (   (   ))     '-'.-(''  ) 
)  )  ( )       --'       '- __.'         :(      )) 
.-'  (_.'          .')                    '(    )  ))
                  (_  )                     ' __.:'          
                                        
--..,___.--,--'',---..-.--+--.,,-,,..._.--..-._.-a:f--.


`
	letterÑ = `
     Ñandú / Rhea
          _
       -=(')
         ;;
        //
       //
      : '.---.__
      |  --_-___)
      \.____,'/
         \  \
       ___\  \
      (       \
               \     jrei
               /

`
	letterO = `
Oveja / Sheep

           __
        ,''''--'''  ''-''-.
      ,'            ,-- ,-'.
     (//            '"'| 'a \
       |    ';         |--._/
       \    _;-._,    /
        \__/\\   \__,'
         ||  ''   \|\\
         \\        \\''
          ''        ''
`

	letterP = `
  Pirata / Pirate
                  _____
              .-" .-. "-.
            _/ '=(0.0)=' \_
          /\   .='|m|'=.   \\
          \________________ /
      .--.__///\'-,__~\\\\~\
     / /6|__\// a (__)-\\\\
     \ \/--\((   ._\   ,)))
     /  \\  ))\  -==-  (O)(
    /    )\((((\   .  /)))))
   /  _.' /  __(\~~~~\)__
  //"\\,-'-"\   \~~~~\\~~\"-.
 //  /\"              \      \
//
  `
	letterQ = `
Queso / Cheese

          ___
        .'o O'-._
       / O o_.-'|
      /O_.-'  O |
      | o   o .-'
jgs   |o O_.-'
      '--'

`
	letterR = `
  Rana / Frog

           .--._.--.
          ( O     O )
          /   . .   \
         .\._______.'.
        /(           )\
      _/  \  \   /  /  \_
   .~   \  \  \ /  /  '   ~.
  {    -.   \  V  /   .-    }
_ _\.    \  |  |  |  /    .'_ _
>_       _} |  |  | {_       _<
 /. - ~ ,_-'  .^.  \-_, ~ - .\
         '-'|/   \|\-\

`
	letterS = `
  Sandwich
                    _.---._
                _.-~       ~-._
            _.-~               ~-._
        _.-~                       ~---._
    _.-~                                 ~\
 .-~                                    _.;
 :-._                               _.-~ ./
 }-._~-._                   _..__.-~ _.-~)
 \-._~-._~-._              / .__..--~_.-~
     ~-._~-._\.        _.-~_/ _..--~~
         ~-. \\--...--~_.-~/~~
            \.\--...--~_.-~
              ~-..----~

  `
	letterT = `
	Triceratops
                       _. - ~ ~ ~ - .
   ..       __.    .-~               ~-.
   ((\     /   \}.~                     \.
    \\\   {     }               /     \   \
(\   \\~~^      }              |       }   \
 \\.-~ -@~      }  ,-.         |       )    \
 (___     ) _}   (    :        |    / /      \.
  \----._-~.     _\ \ |_       \   / /- _      \.
         ~~----~~  \ \| ~~--~~~(  + /     ~-.     ~- _
                   /  /         \  \          ~- . _ _~_-_.
                __/  /          _\  )
              .<___.'         .<___/`
	letterTOption2 = `
	Toro / Bull
 ⠀⠀⠀⠀⠀⠀⠀⢀⣠⣤⣴⣶⣶⡦⠤⠤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢰⣿⣛⣉⣉⣉⣩⣭⣥⣤⣤⣤⡤⢀⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠈⠀⠉⠉⠁⠀⠀⠀⠀⠀⠀⠈⠉⢢⠆⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠤⠤⠤⠄⢀⣀⣀⣀⡘⡄⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⡠⠐⠁⠀⠀⠀⠀⡀⠀⠀⢴⣶⣧⡀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⡠⠊⠀⠀⠀⠀⠀⠀⠀⠹⡄⠀⠨⣿⣿⣷⡄⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⡸⠁⠀⠀⠀⠀⠀⠀⢰⠀⠀⠙⣤⣶⣿⣿⣿⣿⡄⠀⠀⠀⠀
⠀⠀⠀⠀⠀⡐⠁⠀⠀⠀⠀⠀⡠⣴⠾⣷⡆⠀⢿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣧⣴⡄⢻⣿⣿⣿⣿⣿⠀⠀⠀⠀
⠀⠀⠀⠀⢸⠀⠀⢠⠀⠀⠀⠀⠀⠀⠈⠉⢉⠿⢿⣆⢿⣿⣿⣿⣿⡀⠀⠀⠀
⠀⠀⠀⠀⠎⠀⠀⣿⡄⠀⠀⠀⠀⠀⠀⠘⠋⢛⣟⠛⠃⠙⠻⠿⣿⡇⠀⠀⠀
⠀⠀⠀⢸⡄⠀⠀⡘⠋⠉⡀⢠⣾⡰⢶⣶⡖⠁⣤⣳⣿⣶⢶⣶⡌⠳⠤⣀⣀
⠀⠀⠀⢸⢠⠀⢀⣿⣿⣶⣿⣿⣿⠇⠀⠁⣷⣄⣈⣙⣛⣿⣿⣿⡲⡒⠒⠒⠊
⠀⠀⠀⠀⣿⣾⣿⣿⣿⣿⣿⣿⡟⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣟⣿⣶⡄⠀⠀
⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀⠐⣿⠿⣿⣿⣿⣿⡿⠋⠀⠙⣿⡇⠀⠀
⠀⠀⠀⣿⣿⡿⠁⠸⣿⣿⣿⣿⣿⣦⠸⠋⢸⣿⣿⣿⡿⠁⠀⠀⠀⢸⣷⡀⠀
⠀⠀⠀⣻⣿⡇⠀⠀⠀⣹⣿⡿⢻⣿⢠⡀⠸⣿⣿⣿⣧⠀⠀⠀⠀⠘⣿⣧⠀
⠀⠀⢠⠉⣿⠇⠀⠀⢰⠋⣿⣰⣁⡟⠀⠁⢼⣿⡿⠿⠏⠀⠀⠀⠀⠀⠋⠟⠀
⠀⠀⢰⣿⠋⠀⠀⠀⢀⣿⡏⠛⠐⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣾⡇⠀⠀⠀⢀⠎⢹⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠜⢹⡇⠀⠀⠀⠾⣶⡾⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠮⣿⠿⠁⠀⠀  `
	letterTOption3 = `
  Torta / Cake
               (        (
             ( )      ( )          (
      (       Y        Y          ( )
     ( )     |"|      |"|          Y
      Y      | |      | |         |"|
     |"|     | |.-----| |---.___  | |
     | |  .--| |,~~~~~| |~~~,,,,'-| |
     | |-,,~~'-'___   '-'       ~~| |._
    .| |~       // ___            '-',,'.
   /,'-'     <_// // _  __             ~,\
  / ;     ,-,     \\_> <<_______________;_)
  | ;    {(_)} _,      . |================|
  | '-._ ~~,,,           | ,,             |
  |     '-.__ ~~~~~~~~~~~|________________|   
  |\         \'----------|
  | '=._                 |
  :     '=.__            |
   \         \'==========|
snd '-._                 |
        '-.__            |
            \'----------|`
	letterU = `
Unicornio / Unicorn

          .
                 /'
                //
            .  //
            |\//7
           /' " \
          .   . .
          | (    \     '._
          |  '._  '    '. '
          /    \'-'_---. ) )
         .              :.'
         |               \
         | .    .   .     .
         ' .    |  |      |
          \^   /_-':     /
          / | |    '\  .'
         / /| |     \\  |
         \ \( )     // /
          \ | |    // /
           L! !   // / Monoceros'95
            [_]  L[_|  R.B.Cleary

`

	letterV = `
  Vaca / Cow
          .        .
           \'.____.'/
          __'-.  .-'__                         .--.
          '_i:'oo':i_'---...____...----i"""-.-'.-"\\
            /._  _.\       :       /   '._   ;/    ;'-._
           (  o  o  )       '-.__.'       '. '.     '-."
            '-.__.-' _.--.                 '-.:
             : '-'  /     ;   _..--,  /       ;
             :      '-._.-'  ;     ; :       :
              :  \      .'    '-._.' :      /
               \  :    /    ____....--\    :
                '._\  :"""""    '.     !.   :
                 : |: :           'www'| \ '|
                 | || |              : |  | :
                 | || |             .' !  | |
                .' !| |            /__I   | |
               /__I.' !                  .' !
                  /__I                  /__I `
	letterW = `
  W.C. / Toilet

  _____
  |   D
  |   |
  |   |
  \___|            _
    ||  _______  -( (-
    |_'(-------)  '-'
       |       /
_____,-\__..__|_____Pr59

`
	letterX = `
Xilófono / Xylophone
⠀⢀⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢸⣿⠋⢻⣿⠀⣿⣿⠿⣿⡇⢀⣤⣤⣤⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢸⣿⣶⣾⣿⠀⣿⣧⣀⣼⡇⢸⣿⠋⠙⣿⡇⢸⣿⠿⣿⣿⠀⣤⣤⣤⣤⡄⠀
⠀⢸⣿⣿⣿⣿⠀⣿⣿⣿⣿⡇⢸⣿⣷⣾⣿⡇⢸⣧⣀⣼⣿⠀⣿⡏⠉⣿⡇⠀
⠀⢸⣿⣿⣿⣿⠀⣿⣿⣿⣿⡇⢸⣿⣿⣿⣿⡇⢸⣿⣿⣿⣿⠀⣿⣿⣾⣿⡇⠀
⠀⢸⣿⣿⣿⣿⠀⣿⣿⣿⣿⡇⢸⣿⣿⣿⣿⡇⢸⡿⠛⢿⣿⠀⣿⡏⠈⣿⡇⠀
⠀⢸⣿⣿⣿⣿⠀⣿⡿⠛⢿⡇⢸⣿⡁⢈⣿⡇⢸⣷⣤⣾⣿⠀⠿⠿⠿⠿⠇⠀
⠀⢸⣿⡀⣸⣿⠀⣿⣷⣤⣾⡇⠘⠿⠿⠿⠿⠃⠈⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀
⠀⠘⠛⠛⠛⠛⠀⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣄⡀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣤⣴⠀⠻⣿⠗⠀⠀⠀⠀
⠀⠀⠀⠀⠛⠿⠶⠶⠶⣶⣤⣤⣤⣴⣶⣾⣟⡛⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣀⣀⣤⣤⡶⠶⠟⠛⠋⠉⠉⠉⠉⠉⠙⠛⠛⠛⠻⠷⠶⠶⠀⣾⣿⡆⠀⠀
⠀⠀⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠀

`
	letterY = `
  Yoga

                                    ..
                                                .-'  \
                                              .':...::L
                                            .':...::::|
                                           /:::.:=:::::L
                                         .':::...:./:::|
                                        /:::..::::/.\:::L
                                       /:::.:....':::L::|
                                    .-'::::...:/d8888b::|
                               ..dMP=:::::...:'d88888Nb.|
                         ..odMMMP.:.'::=:...'d888888888::L
                       .dMMMMMP..: .:':::.d888888888888I:|
                     .dMMMMMM@b: ' ...:::d8888888888888::|
                    dMMMMMMMMNM.. ..:::d88888888888888P|||
                   dMMMMMMMMMMMboooodP" \?888888888P'\\||
                 .dMMMMMMMMMMMMMMMP'        '"""''  /  |\
                 dMMMMMMMMMMMMMMP'                .'   |
  .mggm..        MMMMMMMMMMMMMP'                .'     |
.dMMMMMMMNNb,    ?MMMMMMMMMMMM|                 |:.. \ |
.MMMMMMMMMMMMMb, .MMMMMMMMMMM(                  \-::   |
.MMMMMMMMMMMMMMMMmdMMMMMMMMMMM.                   \::  |
 ?MMMMMMMMMMMMMMMMMMMMMMMMMMMMb                    |:..|
 \?MMMMMMMMMMMMMMMMMMMMMMMMMMMMbo.,               .MMMMb.
  \MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMb,             dMMMMM:
   ?MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMb           .MMMMMMb
    ?MMMMMMb ?MMMMMMMMMMMMMMMMMMMMMMMMb          .MMMMMMM
     ?MMMMMM  \?MMMMMMMMMMMMMMMMMMMMMMP          dMMMMMMM,
     \?MMMMM     \?WMMMMMMMMMMMMMMMMMMbo,       .MMMMMMMM|
      \MMMMM.         \?MMMMMMMMMMMMMMMMMbo,.   .MMMMMMMMb
       ?MMMMb            \?MMMMMMMMMMMMMMMMMMb, dMMMMMMMMM
       :::::|               \?MMMMMMMMMMMMMMMMMMNMMMMMMMMM
       ::::.|                 \?MMMMMMMMMMMMMMMMMMMMMMMNMN.
    .-:::::.\                    \?MMMMMMMMMMMMMMMMMMMMMHM'
 _.:::::::   |                      \?MMMMMMMMMMMMMMMMMMMM
(_.__________/                         \?MMMM#MMMMMMMMMMMP
                                            \"?MMMMMMMMP'
                                                 \?MMMP

`
	letterZ = `
Zorro / Fox

 ^...^
<_* *_>   
  \_/
        cp97

`
)

func handleUserInput(inputReader *strings.Reader, outputWriter *strings.Builder) {
	fmt.Fprintln(outputWriter, "Había una vez, en un lugar muy muy lejano. Había un ... / Once upon a time, in a land far far away. There was a ...")

	var inputLetter string

	asciiOptions := map[string][]string{
		"A": {letterA},
		"a": {letterA},
		"B": {letterB, letterBOption2, letterBOption3},
		"b": {letterB, letterBOption2, letterBOption3},
		"C": {letterC},
		"c": {letterC},
		"D": {letterD},
		"d": {letterD},
		"E": {letterE},
		"e": {letterE},
		"F": {letterF},
		"f": {letterF},
		"G": {letterG},
		"g": {letterG},
		"H": {letterH},
		"h": {letterH},
		"I": {letterI},
		"i": {letterI},
		"J": {letterJ},
		"j": {letterJ},
		"K": {letterK},
		"k": {letterK},
		"L": {letterL},
		"l": {letterL},
		"M": {letterM},
		"m": {letterM},
		"N": {letterN},
		"n": {letterN},
		"Ñ": {letterÑ},
		"ñ": {letterÑ},
		"O": {letterO},
		"o": {letterO},
		"P": {letterP},
		"p": {letterP},
		"Q": {letterQ},
		"q": {letterQ},
		"R": {letterR},
		"r": {letterR},
		"S": {letterS},
		"s": {letterS},
		"T": {letterT, letterTOption2, letterTOption3},
		"t": {letterT, letterTOption2, letterTOption3},
		"U": {letterU},
		"u": {letterU},
		"V": {letterV},
		"v": {letterV},
		"W": {letterW},
		"w": {letterW},
		"X": {letterX},
		"x": {letterX},
		"Y": {letterY},
		"y": {letterY},
		"Z": {letterZ},
		"z": {letterZ},
	}

	seed := time.Now().UnixNano()

	r := rand.New(rand.NewSource(seed))

	for {
		fmt.Fprint(outputWriter, "Ingresa una letra o 'fin' para terminar: \n / Enter a letter or type 'end' to finish: ")
		fmt.Fscanln(inputReader, &inputLetter)

		if (inputLetter == "fin") || (inputLetter == "end") {
			break
		}

		options, found := asciiOptions[inputLetter]

		if !found {
			fmt.Println("La letra ingresada no tiene una imagen ASCII asociada. \n / The entered letter does not have an associated ASCII image.")
			continue
		}

		selectedASCII := options[r.Intn(len(options))]
		fmt.Fprintln(outputWriter, selectedASCII)

		fmt.Fprintln(outputWriter, "Ingresa otra letra para ver con qué, quién o qué lugar se encuentra nuestro personaje (o escribe 'fin' para terminar): \n / Enter another letter to see with what, who, or what place our character encounters (or type 'end' to finish): ")

	}

}

func main() {
	handleUserInput()
	fmt.Println("¡Hasta luego! \n Good bye!")

}
