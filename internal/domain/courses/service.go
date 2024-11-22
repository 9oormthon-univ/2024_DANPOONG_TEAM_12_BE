package courses

import (
	"github.com/9oormthon-univ/2024_DANPOONG_TEAM_12_BE/internal/types"
)

type CoursesService struct {
	coursesRepository *CoursesRepository
	types.AIService
	types.RegionsService
}

func SetCoursesService(repository *CoursesRepository) types.CoursesService {
	r := &CoursesService{
		coursesRepository: repository,
	}
	return r
}

func (c *CoursesService) RecommendCourses(req *types.RecommendCourseReq) (*types.Course, error) {
	aiCourses, err := c.AIService.RecommendCourses(req)
	if err != nil {
		return nil, err
	}

	result := &types.Course{
		Name: req.AreaCode,
		Plans: []types.Plan{
			{
				DayNumber: "1",
				Places:    []types.Place{},
			},
		},
	}

	for _, aiCourse := range aiCourses {
		contentID := aiCourse.ContentID

		detailData, err := c.RegionsService.GetDetailCommon(contentID)
		if err != nil {
			return nil, err
		}

		// course, place, plan 구조로 응답 생성
		// 예시로 1일차로 설정
		result.Plans[0].Places = append(result.Plans[0].Places, types.Place{
			Name:        detailData.Title,
			Description: detailData.Overview,
			Address:     detailData.Addr1,
			StartTime:   aiCourse.StartTime,
			EndTime:     aiCourse.EndTime,
			Type:        aiCourse.Type,
			MapX:        detailData.MapX,
			MapY:        detailData.MapY,
			ImageURL:    detailData.FirstImage2, // 썸네일
		})

	}
	return result, nil
}

func (c *CoursesService) InjectAIService(service types.AIService) {
	c.AIService = service
}

func (c *CoursesService) InjectRegionService(service types.RegionsService) {
	c.RegionsService = service
}
