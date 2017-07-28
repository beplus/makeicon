package image_helper

var SplashResized = []ResizeImage{

	// Tablet (iPad) non-retina
	{
		"/" + Ios + "/", "Default-Portrait-1024h_768x1024.png", 768, 1024,
	},
	{
		"/" + Ios + "/", "Default-Landscape-1024h_1024x768.png", 1024, 768,
	},
	// iPad retina
	{
		"/" + Ios + "/", "Default-Portrait-1024h@2x_1536x2048.png", 1536, 2048,
	},
	{
		"/" + Ios + "/", "Default-Landscape-1024h@2x_2048x1536.png", 2048, 1536,
	},
	// Handheld (iPhone, iPod) non-retina
	{
		"/" + Ios + "/", "Default-Portrait-480h_320x480", 320, 480,
	},
	{
		"/" + Ios + "/", "Default-Landscape-480h_320x480", 480, 320,
	},
	// Handheld (iPhone, iPod) retina
	{
		"/" + Ios + "/", "Default-Portrait-480h@2x_640x960", 640, 960,
	},
	{
		"/" + Ios + "/", "Default-Landscape-480h@2x_960x640", 960, 640,
	},
	// iPhone 5 Retina
	{
		"/" + Ios + "/", "Default-Portrait-568h@2x_640x1136.png", 640, 1136,
	},
	{
		"/" + Ios + "/", "Default-Landscape-568h@2x_1136x640.png", 1136, 640,
	},
	// iPhone 6 (2x)
	{
		"/" + Ios + "/", "Default-Portrait-667h@2x_750x1334.png", 750, 1334,
	},
	{
		"/" + Ios + "/", "Default-Landscape-667h@2x_1334x750.png", 1334, 750,
	},
	//iPhone 6 Plus (3x)
	{
		"/" + Ios + "/", "Default-Portrait-736h@3x_1242x2208.png", 1242, 2208,
	},
	{
		"/" + Ios + "/", "Default-Landscape-736h@3x_2208x1242.png", 2208, 1242,
	},
	// Android
	// LDPI
	{
		"/" + Android + "/drawable-port-ldpi/", "default", 200, 320,
	},
	{
		"/" + Android + "/drawable-long-ldpi/", "default", 320, 200,
	},
	// MDPI
	{
		"/" + Android + "/drawable-port-mdpi/", "default", 320, 480,
	},
	{
		"/" + Android + "/drawable-long-mdpi/", "default", 480, 320,
	},
	// HDPI
	{
		"/" + Android + "/drawable-port-hdpi/", "default", 480, 800,
	},
	{
		"/" + Android + "/drawable-long-hdpi/", "default", 800, 480,
	},
	// XHDPI
	{
		"/" + Android + "/drawable-port-xhdpi/", "default", 720, 1280,
	},
	{
		"/" + Android + "/drawable-long-xhdpi/", "default", 1280, 720,
	},
	// XXHDPI
	{
		"/" + Android + "/drawable-port-xxhdpi/", "default", 960, 1600,
	},
	{
		"/" + Android + "/drawable-long-xxhdpi/", "default", 1600, 960,
	},
	// XXXHDPI
	{
		"/" + Android + "/drawable-port-xxxhdpi/", "default", 1280, 1920,
	},
	{
		"/" + Android + "/drawable-long-xxxhdpi/", "default", 1920, 1280,
	},
}
