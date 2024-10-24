// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "geocity/static/pages/templates"

func GamingPage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"p-10 mx-32 bg-zinc-900 bg-cover\" style=\"background-image: url(&#39;/media/mvc_avengers.jpg&#39;);\"><div class=\"flex flex-row justify-between\"><div><a class=\"text-white\" href=\"/\">(Break me to go back home) <img src=\"media/smash_ball.gif\" height=\"50\" width=\"50\"></a></div><div><p class=\"text-white\">Enjoy some Alex Certified Music<p class=\"text-xs italic text-white\">(haha get it, cause my name is alex and this is the \"Alex & Ken's Theme\" from SF3), no? Sorry...</p></p><audio controls autoplay loop><source src=\"media/Alex &amp; Ken&#39;s Theme - Jazzy NYC &#39;99 [SF III.3] [ ezmp3.cc ].mp3\" type=\"audio/mp3\"> Your browser does not support the audio element!</audio></div></div><img src=\"media/RyuGIF2.gif\" class=\"absolute\"><div class=\"flip\"><img src=\"media/RyuGIF2.gif\" class=\"absolute\"></div><div class=\"p-2 bg-zinc-700 border-8 border-double border-green-400 text-white\"><h1 class=\"text-5xl text-center py-5 font-gaming-light rainbow\">Welcome to the Gamer Lounge!</h1><p class=\"text-center text-xl my-10\">Nice job breaking the smash ball! Welcome to the venue!</p><p class=\"text-center text-xl my-3\">I'm pretty into gaming <i>(shocker)</i>. I've been playing since a  very young age with my first console being a ps1. Below are some of my fav games below (no particular order besides portal being #1):</p><div class=\"flex flex-col\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var3 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"content-center p-3 leading-relaxed text-left\">This game is an absolute gem. I was presented this on a beautiful  Christmas back in December of 2007 when I was 9 years old.  My dad had bought Valve’s amazing <a class=\"text-orange-400\" href=\"https://en.wikipedia.org/wiki/The_Orange_Box\">\"The Orange Box\" </a>Collection,  and while I now speculate my Dad had bought it so that  he could selfishly try out Team Fortress 2 as well as the Half-Life series,  he saw that Portal was a lower maturity rating of the two and would allow me to play that.  Besides the nostalgia aspect, Portal is just an amazing game in general.  The unsettling atmosphere of the game, you’re constantly under observation with cameras  tracking your every move within the chambers, the blurred windowed rooms above where you  can’t quite make out what’s on the other side looking in.  The music which would bounce between synthetic puzzle tracks,  to what I can only describe as a feeling of hauntingly vacant ambient music.  The only human voice you hear along the way throughout the game turns out to be a murderous AI who attempts to cook you alive.  The gameplay speaks for itself, with what feels like such a simple concept that builds on top of itself and continues to show you  more and more creative ways to utilize the concept of portals to solve the chambers and challenges within Aperture Science.</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = templates.GameCard("Portal", "portal_box.jpg").Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
				if !templ_7745c5c3_IsBuffer {
					defer func() {
						templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err == nil {
							templ_7745c5c3_Err = templ_7745c5c3_BufErr
						}
					}()
				}
				ctx = templ.InitializeContext(ctx)
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"content-center p-3 leading-relaxed text-left\">This was my very first game I ever played on PC. I absolutey love this game series, and the remaster that recently came out for it I was in love with. The visual asthetic of this game is something of its own. Something about seeing a simple trailed color light of your space fleet  as you command them around is just so satisfyting. My absolute favorite ship in this game was the Salvage Corvette, While my Dad and I do not see eye to eye, I can’t deny that some of my fondest memories as a kid  were watching / playing this with him. The encouragement he gave me for completing the resource managment  and strategy aspects of the game as well as the game world I would say domino my passion for SciFi, tech, and general interest in problem solving.</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return templ_7745c5c3_Err
			})
			templ_7745c5c3_Err = templates.GameCard("Homeworld", "homeworld_box.jpg").Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Pokemon Emerald", "emerald_box.jpg").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Dragon Ball Z: Budokai Tenkaichi 3", "dbzbt3_box.jpg").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Team Fortress 2", "tf2_box.jpg").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Super Smash Bros. Ultimate", "ssbu_box.jpg").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Dragon Ball FighterZ", "dbfz_box.jpg").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Halo Reach", "reach_box.png").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Minecraft", "minecraft_box.png").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Dark Souls 3", "ds3_box.jpg").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templates.GameCard("Metal Gear Solid V: The Phantom Pain", "mgsv_box.png").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></body>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = templates.Base(templates.DefaultBase()).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
