Each game must provide a persistent 64-bit ID for each asset that can be traded to another user. At a minimum this asset ID must be unique within the context that contains the asset ID.
Steam reserves the top two bits of each context ID for its own use. Context IDs returned by an asset server must be zero in the top two bits.
All contexts defined by the asset server must also have a non-zero context ID. APIs that return context IDs must not specify a value of zero.
Contexts provide a way to break up items into different categories, what kind of categories depend on the implementation.

Ссылка на получение изображения на основ еинформации о предмете
---------------------------------------------------------------

https://steamcommunity-a.akamaihd.net/economy/image/class/730/469438644

App ID (appid)
--------------

Целочисленный идентфикатор игры (730 - csgo)

Context ID (contextid)
----------------------

Способ организации/категоризации
элементов(items)/имущества(assets)/валют(currency). Это просто число, но
документация стима указывает на способ сделать нечто вроде директорий с помощью
разделения числа на битовые последовательности (bits ranges) и использования
каждой битовой последовательности (bit range) для обозначения чего-то другого

Class ID (classid)
------------------

Это ID который идентифицирует класс элемента (item) свойства которого все
одинаковый для всех элементов (items) с таким же class ID

Instance ID (instanceid)
------------------------

Это ID которые описывает экземпляр элемента (item) который наследует свойства
из класса, class ID которого указан в данном поле. По всей видимости это ссылка
на class ID существующего товара

Asset ID (assetid)
------------------

Уникальный идентификатор имущества(asset) (элемента(item)). Это имущество будет
иметь свой собственный id, и так же class ID и instance ID для указания на
фактическое представление элемента. Объекты имущества могут так же иметь
количественные свойства для указания как много элементов с <classID_instanceID>
пользователь может стакать (иными словами - максимальная вместимость ячейки для
всех предметов с таким же class ID и instanceId). Asset ID может так же
измениться когда элемент передается, хотя class ID и instance ID должен остаться
тот же

Currency ID (currencyid)
------------------------

Это ID который идентифицирует кастомную валюту созданную и реализованную игрой.
Этот идентификатор остается одинаковым везде, в отличие от asset ID.

Definition ID (defid)
---------------------

Идентификатор, уникальный в масштабе приложения который стим используется внутри
для соотношения элементов (items) к атрибутам в файле схемы (??)(schema file).
Наверное это что то про инсталяцию стима на компе


Каждый объект в экономике стима может быть однозначно идентифицирован
комбинацией трех чисел: App ID, Context ID, и Asset ID


Item class_instance Info
	icon_url			If the icon_url is not a url then prefix this url to the value of icon_url "https://steamcommunity-a.akamaihd.net/economy/image/"
						And to get different sizes one can add a "/" to the end of the url and attach the dimensions the app wishes to retrieve the image as.
						This will resize the image to fit within the given dimensions maintaining the aspect ratio of the original image.
						If one wants to force dimensions for an image just add an f after the dimension, /300fx50
							This will still resize the image to fit within the given dimensions maintaining the aspect ratio, but it will be centered within the given dimensions and the complete image will have given dimensions.
	icon_drag_url		The icon image that should be used when the item is being dragged by a mouse cursor or touch event. If value is an empty string then just use the value of icon_url
	tradeable			Whether or not the asset / currency is tradeable
	marketable			Whether or not the asset / currency can be sold in the steam marktplace
	name				The name of the asset / currency
	name_color			The color of the text the name should be displayed in and also the color of accents such as the border for the item.
	background_color	The background color the item / icon should be displayed on.
	market_name			The name of the item when viewed in the marketplace.
	type				The type of item, this value is suitable for presentation.
	fraudwarnings		A list of fraud warning objects that should be displayed to a user to alert the of potential malicious activity.
	descriptions		An array of description objects that should be displayed with the item when its details are viewed.
						Description objects normally have 2 properties "type", and "value" but may have any of the properties listed below, or not listed.
						type
							values:
							html - Represents an HTML block element

						value
							The value to use as type

						color
							if type = html then the color of the text for that html object

						app_data
							Data that is relevant to the game that pertains to this description.
	actions				A list of action objects that can be used to take an action on an item.
						An action object normally has 2 properties, "name" and "link"
						name - display name of the action eg "Inspect in Game"
						link - The link to execute to perform the action, this may not even be http, it could be any valid uri protocol registered on the local computer.
	market_actions		A list of action objects that can be used to take an action on the item when viewed from the steam marketplace. Action objects are the same as described in the "actions" section.
	tags				A list of tag objects that allow for advanced filtering when viewing items in an inventory setting.
						A tag object normally has 4 properties: "internal_name", "name", "category", and "category_name"
						internal_name	Name used internally by the game and filtering functions of the inventory.
						name			Name to be displayed to the user as a filter.
						category		The internal name of the category the tag filter should be listed under
						category_name	The name to be displayed to the user as the category.
						color			The color to display the text as.
						Tag objects can be thought of as a list of key value pairs where item.category = internal_name

	Any of the above listed properties can be overwritten by instance data allowing for unique values per instance.
