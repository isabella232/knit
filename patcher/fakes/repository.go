package fakes

type Repository struct {
	ConfigureCommitterCall struct {
		Count   int
		Returns struct {
			Error error
		}
	}

	CheckoutCall struct {
		Receives struct {
			Ref string
		}
		Returns struct {
			Error error
		}
	}

	CleanSubmodulesCall struct {
		Count   int
		Returns struct {
			Error error
		}
	}

	ApplyPatchCall struct {
		Receives struct {
			Patches []string
		}
		Returns struct {
			Error error
		}
	}

	BumpSubmoduleCall struct {
		Receives struct {
			Submodules map[string]string
		}
		Returns struct {
			Error error
		}
	}

	PatchSubmoduleCall struct {
		Receives struct {
			Paths   []string
			Patches []string
		}
		Returns struct {
			Error error
		}
	}

	CheckoutBranchCall struct {
		Receives struct {
			Name string
		}
		Returns struct {
			Error error
		}
	}
}

func (r *Repository) ConfigureCommitter() error {
	r.ConfigureCommitterCall.Count++

	return r.ConfigureCommitterCall.Returns.Error
}

func (r *Repository) Checkout(checkoutRef string) error {
	r.CheckoutCall.Receives.Ref = checkoutRef

	return r.CheckoutCall.Returns.Error
}

func (r *Repository) CleanSubmodules() error {
	r.CleanSubmodulesCall.Count++

	return r.CleanSubmodulesCall.Returns.Error
}

func (r *Repository) ApplyPatch(patch string) error {
	r.ApplyPatchCall.Receives.Patches = append(r.ApplyPatchCall.Receives.Patches, patch)

	return r.ApplyPatchCall.Returns.Error
}

func (r *Repository) BumpSubmodule(patchPath, sha string) error {
	if len(r.BumpSubmoduleCall.Receives.Submodules) == 0 {
		r.BumpSubmoduleCall.Receives.Submodules = make(map[string]string)
	}

	r.BumpSubmoduleCall.Receives.Submodules[patchPath] = sha

	return r.BumpSubmoduleCall.Returns.Error
}

func (r *Repository) PatchSubmodule(relativePath, fullPathToPatch string) error {
	r.PatchSubmoduleCall.Receives.Paths = append(r.PatchSubmoduleCall.Receives.Paths, relativePath)
	r.PatchSubmoduleCall.Receives.Patches = append(r.PatchSubmoduleCall.Receives.Patches, fullPathToPatch)

	return r.PatchSubmoduleCall.Returns.Error
}

func (r *Repository) CheckoutBranch(name string) error {
	r.CheckoutBranchCall.Receives.Name = name

	return r.CheckoutBranchCall.Returns.Error
}
