package main

import (
	"testing"

	pluginv1 "github.com/Silo-Server/silo-plugin-sdk/pkg/pluginproto/silo/plugin/v1"
	publicmanifest "github.com/Silo-Server/silo-plugin-sdk/pkg/pluginsdk/manifest"
)

func TestAdminFormLayout(t *testing.T) {
	m, err := publicmanifest.LoadWithChecksum(manifestJSON, version)
	if err != nil {
		t.Fatalf("LoadWithChecksum: %v", err)
	}
	af := m.GetCapabilities()[0].GetConfigSchema()[0].GetAdminForm()

	var library, anime *pluginv1.AdminFormSection
	for _, s := range af.GetSections() {
		switch s.GetKey() {
		case "library":
			library = s
		case "anime":
			anime = s
		}
	}
	if library == nil || anime == nil {
		t.Fatalf("expected both library and anime sections, got %+v", af.GetSections())
	}

	// Library: collapsible, collapsed by default, and no longer owns the gate toggle.
	if !library.GetCollapsible() || !library.GetCollapsedDefault() {
		t.Errorf("library: collapsible=%v collapsed_default=%v, want both true", library.GetCollapsible(), library.GetCollapsedDefault())
	}
	for _, k := range library.GetFieldKeys() {
		if k == "anime_enabled" {
			t.Errorf("anime_enabled must not be in the library section")
		}
	}

	// Anime: always visible (no section-level show_when, not collapsible) and
	// gated by anime_enabled as its first field.
	if anime.GetCollapsible() {
		t.Errorf("anime section must not be collapsible (the gate toggle must stay visible)")
	}
	if len(anime.GetShowWhen()) != 0 {
		t.Errorf("anime section must not carry a section-level show_when, got %+v", anime.GetShowWhen())
	}
	keys := anime.GetFieldKeys()
	if len(keys) == 0 || keys[0] != "anime_enabled" {
		t.Errorf("anime section must list anime_enabled first, got %v", keys)
	}
}

func TestEmbeddedManifestLoads(t *testing.T) {
	m, err := publicmanifest.LoadWithChecksum(manifestJSON, version)
	if err != nil {
		t.Fatalf("LoadWithChecksum: %v", err)
	}
	if m.GetPluginId() != "silo.requests.arr" {
		t.Fatalf("plugin_id = %q", m.GetPluginId())
	}
	if len(m.GetCapabilities()) != 1 || m.GetCapabilities()[0].GetType() != "request_router.v1" {
		t.Fatalf("expected one request_router.v1 capability, got %+v", m.GetCapabilities())
	}
}
