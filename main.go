package main

import (
  "github.com/fatih/color"

  "bufio"
  "fmt"
  "os"
)

func main() {
  fmt.Println("\nlorem v1.0 - fabricioh")
  color.HiYellow("\n\tPRESS ENTER TO CONTINUE")
  color.HiYellow("\tINPUT 'x' TO QUIT\n\n")

  input := bufio.NewScanner(os.Stdin)

  ipsum := []string{
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Amet est placerat in egestas erat imperdiet sed. Ac feugiat sed lectus vestibulum mattis. Venenatis a condimentum vitae sapien pellentesque habitant morbi. Neque sodales ut etiam sit amet nisl purus in mollis. Enim sit amet venenatis urna. Eu mi bibendum neque egestas congue. Dictumst quisque sagittis purus sit amet volutpat consequat mauris. A erat nam at lectus urna duis convallis convallis tellus. Nibh tortor id aliquet lectus. Tellus elementum sagittis vitae et leo duis ut diam. Viverra nam libero justo laoreet sit amet cursus sit. Et pharetra pharetra massa massa ultricies mi quis hendrerit. Quam viverra orci sagittis eu volutpat. At volutpat diam ut venenatis tellus in.",
    "Ac felis donec et odio pellentesque diam. Vitae sapien pellentesque habitant morbi tristique senectus et netus. Dictum varius duis at consectetur lorem donec massa sapien faucibus. Tellus integer feugiat scelerisque varius morbi enim. Aliquam id diam maecenas ultricies mi eget mauris pharetra. Magna eget est lorem ipsum dolor sit. Venenatis lectus magna fringilla urna. Nulla pellentesque dignissim enim sit amet venenatis urna. Leo urna molestie at elementum eu facilisis sed. Condimentum lacinia quis vel eros donec ac odio tempor orci.",
    "Eleifend donec pretium vulputate sapien nec sagittis aliquam malesuada bibendum. Dictum fusce ut placerat orci nulla. At augue eget arcu dictum varius duis at consectetur lorem. Platea dictumst vestibulum rhoncus est pellentesque elit ullamcorper dignissim. Habitant morbi tristique senectus et netus et malesuada fames ac. Mi tempus imperdiet nulla malesuada pellentesque elit eget. Convallis convallis tellus id interdum velit laoreet id donec ultrices. Odio ut sem nulla pharetra diam sit amet nisl. At consectetur lorem donec massa sapien. Commodo odio aenean sed adipiscing diam donec adipiscing. Neque convallis a cras semper auctor neque vitae tempus. Duis at consectetur lorem donec massa sapien faucibus et molestie. Aliquet enim tortor at auctor urna nunc id cursus metus. Facilisi etiam dignissim diam quis enim lobortis scelerisque. Vitae aliquet nec ullamcorper sit amet. Fermentum iaculis eu non diam phasellus vestibulum lorem sed risus. Viverra suspendisse potenti nullam ac tortor vitae. Ultricies tristique nulla aliquet enim tortor at auctor.",
    "Ultrices mi tempus imperdiet nulla malesuada pellentesque elit. Euismod in pellentesque massa placerat duis ultricies lacus sed. Sem viverra aliquet eget sit amet tellus cras adipiscing enim. Sit amet justo donec enim diam vulputate ut pharetra. Nisi vitae suscipit tellus mauris a. Pellentesque adipiscing commodo elit at imperdiet dui accumsan. Ante metus dictum at tempor commodo ullamcorper a. Purus in mollis nunc sed id semper. Neque sodales ut etiam sit amet nisl purus. Pellentesque diam volutpat commodo sed egestas egestas. Fusce id velit ut tortor. Nulla aliquet porttitor lacus luctus accumsan tortor posuere ac ut.",
    "Non tellus orci ac auctor augue mauris augue neque. Faucibus ornare suspendisse sed nisi lacus sed viverra tellus. Nunc mattis enim ut tellus. Id venenatis a condimentum vitae. Mi sit amet mauris commodo quis imperdiet. Quam elementum pulvinar etiam non. Sed risus pretium quam vulputate dignissim. Consectetur adipiscing elit pellentesque habitant morbi. Nisl pretium fusce id velit ut tortor pretium viverra suspendisse. Pretium quam vulputate dignissim suspendisse in. Malesuada pellentesque elit eget gravida cum sociis natoque penatibus et. Ac odio tempor orci dapibus ultrices in iaculis nunc sed. Accumsan in nisl nisi scelerisque eu. Eget dolor morbi non arcu risus quis varius. Id volutpat lacus laoreet non curabitur. Dui faucibus in ornare quam viverra orci sagittis eu volutpat. Integer malesuada nunc vel risus commodo viverra. Duis at tellus at urna condimentum mattis.",
    "Diam vulputate ut pharetra sit amet aliquam id diam. Purus non enim praesent elementum facilisis leo vel. Vitae congue eu consequat ac felis donec et odio. Erat pellentesque adipiscing commodo elit at imperdiet. Odio aenean sed adipiscing diam donec adipiscing tristique. Mi ipsum faucibus vitae aliquet nec ullamcorper sit. Sed libero enim sed faucibus turpis in eu mi. Fames ac turpis egestas integer eget aliquet nibh praesent tristique. Quam quisque id diam vel quam elementum pulvinar. Porttitor leo a diam sollicitudin tempor id eu.",
    "Aliquam nulla facilisi cras fermentum odio eu feugiat pretium nibh. Aenean et tortor at risus viverra adipiscing. Elit ullamcorper dignissim cras tincidunt lobortis feugiat vivamus at. Bibendum at varius vel pharetra vel turpis nunc eget. Nulla posuere sollicitudin aliquam ultrices sagittis. Lectus arcu bibendum at varius vel. Morbi tristique senectus et netus et. Ut tellus elementum sagittis vitae et leo duis ut. Commodo odio aenean sed adipiscing diam. Nunc scelerisque viverra mauris in aliquam sem. Amet luctus venenatis lectus magna fringilla urna porttitor rhoncus. Est velit egestas dui id ornare arcu. Nunc sed velit dignissim sodales ut. Lorem sed risus ultricies tristique nulla aliquet. Ipsum dolor sit amet consectetur adipiscing. Sit amet consectetur adipiscing elit. Eu turpis egestas pretium aenean pharetra magna.",
    "Arcu dictum varius duis at. Egestas maecenas pharetra convallis posuere morbi leo urna. Viverra nibh cras pulvinar mattis nunc sed blandit. Sit amet risus nullam eget felis eget nunc lobortis. Elit sed vulputate mi sit amet mauris commodo quis imperdiet. Porttitor eget dolor morbi non arcu risus quis. Adipiscing elit duis tristique sollicitudin nibh sit. Nunc vel risus commodo viverra. Scelerisque varius morbi enim nunc faucibus a. Sed pulvinar proin gravida hendrerit lectus.",
    "Enim ut sem viverra aliquet eget sit amet. Posuere sollicitudin aliquam ultrices sagittis. Est placerat in egestas erat imperdiet sed euismod nisi. Platea dictumst vestibulum rhoncus est pellentesque elit. Condimentum id venenatis a condimentum vitae sapien. Pharetra magna ac placerat vestibulum lectus mauris ultrices eros. Sapien nec sagittis aliquam malesuada bibendum arcu vitae elementum. Ac tortor dignissim convallis aenean et. Volutpat commodo sed egestas egestas fringilla phasellus faucibus scelerisque. Integer enim neque volutpat ac tincidunt vitae semper. Eu mi bibendum neque egestas congue quisque egestas diam. Nisi scelerisque eu ultrices vitae auctor eu augue ut. Amet nulla facilisi morbi tempus iaculis urna id.",
    "Congue nisi vitae suscipit tellus. Feugiat scelerisque varius morbi enim nunc faucibus a pellentesque sit. Ipsum dolor sit amet consectetur. Tortor condimentum lacinia quis vel eros donec. Ut tristique et egestas quis. Non blandit massa enim nec dui nunc mattis enim. Nam at lectus urna duis convallis convallis tellus id interdum. Lectus nulla at volutpat diam ut venenatis tellus. Quam lacus suspendisse faucibus interdum posuere lorem ipsum dolor. Morbi non arcu risus quis varius quam quisque id.",
  }

  for _, para := range ipsum {
    fmt.Println(para)
    input.Scan()

    if input.Text() == "x" {
      os.Exit(0)
    }
  }
}
