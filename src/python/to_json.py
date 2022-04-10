import json

import pandas as pd

gamemode = {
    "CQ Large": "ConquestLarge0",
    "CQ Small": "ConquestSmall0",
    "Rush": "RushLarge0",
    "SDM": "SquadDeathMatch0",
    "TDM": "TeamDeathMatch0",
    "TDM CQ": "TeamDeathMatchC0",
    "Assault Large": "ConquestAssaultLarge0",
    "Assault": "ConquestAssaultSmall0",
    "Assault 2": "ConquestAssaultSmall1",
    "CQ Dom": "Domination0",
    "GM": "GunMaster0",
    "Tank Superiority": "TankSuperiority0",
    "SQ Rush": "SquadRush0",
    "CTF": "CaptureTheFlag0"
}

name_translate = {
    "Grand Bazaar": "大型市集",
    "Teheran Highway": "德黑兰公路",
    "Caspian Border": "里海边境",
    "Seine Crossing": " 塞纳河渡口",
    "Operation Firestorm": "火线风暴",
    "Damavand Peak": "德玛峰顶",
    "Noshahr Canals": "诺沙尔运河",
    "Kharg Island": "哈尔克岛",
    "Operation Metro": "地铁行动",
    "Strike at Karkand": "突袭卡肯",
    "Gulf of Oman": "阿曼湾",
    "Sharqi Peninsula": "沙尔吉半岛",
    "Wake Island": "威克岛",
    "Scrapmetal": "废金属",
    "Operation 925": "白领行动",
    "Ziba Tower": "济巴塔",
    "Donya Fortress": "唐亚要塞",
    "Alborz Mountains": "厄尔布尔士山脉",
    "Bandar Desert": "班达尔沙漠",
    "Armored Shield": "装甲护罩",
    "Death Valley": "死亡谷",
    "Markaz Monolith": "玛卡斯巨石",
    "Talah Market": "塔拉市场",
    "Epicenter": "震央",
    "Azadi Palace": "阿扎迪王宫",
    "Operation Riverside": "河岸行动",
    "Nebandan Flats": "内巴丹平原",
    "Kiasar Railroad": "奇亚沙铁路",
    "Sabalan Pipeline": "萨巴兰输油管"
}

def abbrToGM(abbr: str):
    gm_list = abbr.split(',')
    mode = {}
    for i in range(len(gm_list)):
        gm_list[i] = gm_list[i].strip()

        mode[gamemode[gm_list[i]]] = 1

    return mode

def main(input_name: str):
    map_info = pd.read_csv(input_name, sep="\t", header=None)
    map_json = {}

    for index, row in map_info.iterrows():
        support_gm = row[3]
        mode = abbrToGM(support_gm)
        support_gm_ve = row[4]
        if isinstance(support_gm_ve, str):
            support_gm_ve = support_gm_ve.split(',')
            for m in support_gm_ve:
                m = m.strip()
                mode[gamemode[m]] = 2

        tech_name = row[1].strip()
        map_name = row[2].strip()
        map_json[tech_name] = {"name": map_name, "cn_name": name_translate[map_name]}
        map_json[tech_name].update(mode)

    with open("./map.json", "w") as file_obj:
        json.dump(map_json, file_obj)