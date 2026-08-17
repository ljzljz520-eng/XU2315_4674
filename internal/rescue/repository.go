package rescue

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Repository interface {
	Home() HomePage
	Articles() []Article
	BasicEquipment() []EquipmentItem
}

type MemoryRepository struct {
	articles []Article
	groups   []EquipmentGroup
	home     HomePage
}

func NewMemoryRepository() *MemoryRepository {
	var articles []Article
	var groups []EquipmentGroup
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		if err := ctx.Err(); err != nil {
			return err
		}
		articles = []Article{
			{ID: "rope-anchors", Title: "复杂地形的绳索锚点记录", Section: "绳索", Summary: "从双点锚固到边缘保护，整理山地现场的检查顺序。", Keywords: []string{"绳索", "锚点", "保护"}},
			{ID: "radio-check", Title: "低温环境下的通信检查", Section: "通信", Summary: "用简短的通联流程确认队伍位置与备用频道。", Keywords: []string{"通信", "对讲机", "频道"}},
			{ID: "warm-layer", Title: "等待救援时的分层保暖", Section: "保暖", Summary: "在风雪和等待中维持体温的装备搭配与记录方法。", Keywords: []string{"保暖", "分层", "风雪"}},
			{ID: "ridge-nav", Title: "山脊路线的导航复盘", Section: "导航", Summary: "把轨迹、等高线和能见度变化放在同一张复盘表里。", Keywords: []string{"导航", "地图", "轨迹"}},
			{ID: "night-evac", Title: "夜间转运案例复盘", Section: "案例复盘", Summary: "一次夜间转运中的分工、照明和风险沟通记录。", Keywords: []string{"案例", "转运", "照明"}},
		}
		return nil
	})
	group.Go(func() error {
		if err := ctx.Err(); err != nil {
			return err
		}
		groups = []EquipmentGroup{
			{Name: "基础装备", Items: []EquipmentItem{
				{Name: "安全头盔", Category: "个人防护", Purpose: "保护头部并便于安装照明"},
				{Name: "主绳与辅绳", Category: "绳索", Purpose: "建立保护、下降和拖拽系统"},
				{Name: "保暖层", Category: "保暖", Purpose: "在停留和等待期间维持体温"},
				{Name: "对讲机", Category: "通信", Purpose: "保持队伍间的短距通联"},
				{Name: "地图与指北针", Category: "导航", Purpose: "在电子设备不可用时确认方向"},
			}},
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		panic(err)
	}
	return &MemoryRepository{
		articles: articles,
		groups:   groups,
		home: HomePage{
			Title: "山地救援装备笔记",
			SafetyTips: []string{
				"出发前逐件核对装备，现场按队伍分工复述关键动作。",
				"天气和路线变化时及时暂停复核，保持通联简洁清楚。",
				"记录事实与判断依据，方便下一次行动快速参考。",
			},
			Featured:   articles[:3],
			AuthorTeam: []string{"林岚 · 绳索系统", "周野 · 山地通信", "顾宁 · 行动复盘"},
			Sections:   []string{"绳索", "通信", "保暖", "导航", "案例复盘"},
			BasicsPath: "/equipment-basics",
		},
	}
}

func (r *MemoryRepository) Home() HomePage {
	return r.home
}

func (r *MemoryRepository) Articles() []Article {
	return append([]Article(nil), r.articles...)
}

func (r *MemoryRepository) BasicEquipment() []EquipmentItem {
	if len(r.groups) == 0 {
		return nil
	}
	return append([]EquipmentItem(nil), r.groups[0].Items...)
}
