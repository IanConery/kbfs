// Copyright 2016 Keybase Inc. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package libfs

// MetricsFileName is the name of the KBFS metrics file -- it can be
// reached from any KBFS directory.
const MetricsFileName = ".kbfs_metrics"

// ReclaimQuotaFileName is the name of the KBFS quota-reclaiming file
// -- it can be reached anywhere within a top-level folder.
const ReclaimQuotaFileName = ".kbfs_reclaim_quota"

// RekeyFileName is the name of the KBFS rekeying file -- it can be
// reached anywhere within a top-level folder.
const RekeyFileName = ".kbfs_rekey"

// StatusFileName is the name of the KBFS status file -- it can be reached
// anywhere within a top-level folder or inside the Keybase root
const StatusFileName = ".kbfs_status"

// SyncFromServerFileName is the name of the KBFS sync-from-server
// file -- it can be reached anywhere within a top-level folder.
const SyncFromServerFileName = ".kbfs_sync_from_server"

// UnstageFileName is the name of the KBFS unstaging file -- it can be
// reached anywhere within a top-level folder.
const UnstageFileName = ".kbfs_unstage"

// DisableUpdatesFileName is the name of the KBFS update-disabling
// file -- it can be reached anywhere within a top-level folder.
const DisableUpdatesFileName = ".kbfs_disable_updates"

// EnableUpdatesFileName is the name of the KBFS update-enabling
// file -- it can be reached anywhere within a top-level folder.
const EnableUpdatesFileName = ".kbfs_enable_updates"

// ResetCachesFileName is the name of the KBFS unstaging file.
const ResetCachesFileName = ".kbfs_reset_caches"
