// Copyright (C) 2012-2021 Miquel Sabaté Solà 
// This file is licensed under the MIT license.
// See the LICENSE file.

package user_agent

import (
	"regexp"
	"strings"
)

var ie11Regexp = regexp.MustCompile("^rv:(.+)$")

// Browser is a struct containing all the information that we might be
// interested from the browser.
type Browser struct {
	// The name of the browser's engine.
	Engine string

	// The version of the browser's engine.
	EngineVersion string

	// The name of the browser.
	Name string

	// The version of the browser.
	Version string
}

// Extract all the information that we can get from the User-Agent string
// about the browser and update the receiver with this information.
// 浏览器相关的信息
// The function receives just one argument "sections", that contains the
// sections from the User-Agent string after being parsed.
func (p *UserAgent) detectBrowser(sections []section) {
	slen := len(sections)

	if sections[0].name == "opera" {
		p.browser.Name = "opera"
		p.browser.Version = sections[0].version
		p.browser.Engine = "presto"
		if slen > 1 {
			p.browser.EngineVersion = sections[1].version
		}
	} else if sections[0].name == "dalvik" {
		// When Dalvik VM is in use, there is no browser info attached to ua.
		// Although browser is still a Mozilla/5.0 compatible.
		p.mozilla = "5.0"
	} else if slen > 1 {
		engine := sections[1]
		p.browser.Engine = engine.name
		p.browser.EngineVersion = engine.version
		if slen > 2 {
			sectionIndex := 2
			// The version after the engine comment is empty on e.g. Ubuntu
			// platforms so if this is the case, let's use the next in line.
			if sections[2].version == "" && slen > 3 {
				sectionIndex = 3
			}
			p.browser.Version = sections[sectionIndex].version
			if engine.name == "applewebkit" {
				for _, comment := range engine.comment {
					if len(comment) > 5 &&
						(strings.HasPrefix(comment, "googlebot") || strings.HasPrefix(comment, "bingbot")) {
						p.undecided = true
						break
					}
				}
				switch sections[slen-1].name {
				case "edge":
					p.browser.Name = "edge"
					p.browser.Version = sections[slen-1].version
					p.browser.Engine = "edgehtml"
					p.browser.EngineVersion = ""
				case "edg":
					if !p.undecided {
						p.browser.Name = "edge"
						p.browser.Version = sections[slen-1].version
						p.browser.Engine = "applewebkit"
						p.browser.EngineVersion = sections[slen-2].version
					}
				case "opr":
					p.browser.Name = "opera"
					p.browser.Version = sections[slen-1].version
				case "mobile":
					p.browser.Name = "mobile app"
					p.browser.Version = ""
				default:
					switch sections[slen-3].name {
					case "yabrowser":
						p.browser.Name = "yabrowser"
						p.browser.Version = sections[slen-3].version
					case "coc_coc_browser":
						p.browser.Name = "coc coc"
						p.browser.Version = sections[slen-3].version
					default:
						switch sections[slen-2].name {
						case "electron":
							p.browser.Name = "electron"
							p.browser.Version = sections[slen-2].version
						case "duckduckgo":
							p.browser.Name = "duckduckgo"
							p.browser.Version = sections[slen-2].version
						default:
							switch sections[sectionIndex].name {
							case "chrome", "crios":
								p.browser.Name = "chrome"
							case "headlesschrome":
								p.browser.Name = "headless chrome"
							case "chromium":
								p.browser.Name = "chromium"
							case "gsa":
								p.browser.Name = "google app"
							case "fxios":
								p.browser.Name = "firefox"
							default:
								p.browser.Name = "safari"
							}
						}
					}
					// It's possible the google-bot emulates these now
					for _, comment := range engine.comment {
						if len(comment) > 5 &&
							(strings.HasPrefix(comment, "googlebot") || strings.HasPrefix(comment, "bingbot")) {
							p.undecided = true
							break
						}
					}
				}
			} else if engine.name == "gecko" {
				name := sections[2].name
				if name == "mra" && slen > 4 {
					name = sections[4].name
					p.browser.Version = sections[4].version
				}
				p.browser.Name = name
			} else if engine.name == "like" && sections[2].name == "gecko" {
				// This is the new user agent from Internet Explorer 11.
				p.browser.Engine = "trident"
				p.browser.Name = "internet explorer"
				for _, c := range sections[0].comment {
					version := ie11Regexp.FindStringSubmatch(c)
					if len(version) > 0 {
						p.browser.Version = version[1]
						return
					}
				}
				p.browser.Version = ""
			}
		}
	} else if slen == 1 && len(sections[0].comment) > 1 {
		comment := sections[0].comment
		if comment[0] == "compatible" && strings.HasPrefix(comment[1], "msie") {
			p.browser.Engine = "trident"
			p.browser.Name = "internet explorer"
			// The MSIE version may be reported as the compatibility version.
			// For IE 8 through 10, the Trident token is more accurate.
			// http://msdn.microsoft.com/en-us/library/ie/ms537503(v=vs.85).aspx#VerToken
			for _, v := range comment {
				if strings.HasPrefix(v, "trident/") {
					switch v[8:] {
					case "4.0":
						p.browser.Version = "8.0"
					case "5.0":
						p.browser.Version = "9.0"
					case "6.0":
						p.browser.Version = "10.0"
					}
					break
				}
			}
			// If the Trident token is not provided, fall back to MSIE token.
			if p.browser.Version == "" {
				p.browser.Version = strings.TrimSpace(comment[1][4:])
			}
		}
	}
}

// Engine returns two strings. The first string is the name of the engine and the
// second one is the version of the engine.
func (p *UserAgent) Engine() (string, string) {
	return p.browser.Engine, p.browser.EngineVersion
}

// Browser returns two strings. The first string is the name of the browser and the
// second one is the version of the browser.
func (p *UserAgent) Browser() (string, string) {
	return p.browser.Name, p.browser.Version
}
