package arr

import "testing"

func TestRouteTargetsRoutesByIs4K(t *testing.T) {
	instances := []Instance{
		{ID: "hd", Kind: "radarr", Enabled: true, IsDefault: true},
		{ID: "uhd", Kind: "radarr", Enabled: true, IsDefault4K: true, Is4K: true},
	}
	req := Request{MediaType: MediaTypeMovie}
	planned := RouteTargets(req, []RequestedQuality{
		{ID: "1080p", Is4K: false},
		{ID: "2160p", Is4K: true},
	}, instances)
	if len(planned) != 2 {
		t.Fatalf("want 2 targets, got %d", len(planned))
	}
	byQ := map[string]string{}
	for _, p := range planned {
		byQ[p.Quality] = p.Instance.ID
	}
	if byQ["1080p"] != "hd" || byQ["2160p"] != "uhd" {
		t.Fatalf("routing wrong: %+v", byQ)
	}
}
