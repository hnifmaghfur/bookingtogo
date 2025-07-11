@extends('layouts.app')

@section('title', 'Create Customer')

@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            <div class="card">
                <div class="card-header">
                    <h4>Create New Customer</h4>
                </div>
                <div class="card-body">
                    @if(session('success'))
                    <div class="alert alert-success alert-dismissible fade show" role="alert">
                        {{ session('success') }}
                        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
                    </div>
                    @endif

                    @if(session('error'))
                    <div class="alert alert-danger alert-dismissible fade show" role="alert">
                        {{ session('error') }}
                        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
                    </div>
                    @endif

                    {!! Form::open(['route' => 'customer.store', 'method' => 'POST', 'id' => 'customerForm']) !!}

                    <!-- Customer Information -->
                    <div class="row">
                        <div class="col-md-6">
                            <div class="mb-3">
                                {!! Form::label('name', 'Name', ['class' => 'form-label']) !!}
                                <span class="text-danger">*</span>
                                {!! Form::text('name', old('name'), [
                                'class' => 'form-control' . ($errors->has('name') ? ' is-invalid' : ''),
                                'id' => 'name',
                                'required' => true,
                                'placeholder' => 'Enter full name'
                                ]) !!}
                                @error('name')
                                <div class="invalid-feedback">{{ $message }}</div>
                                @enderror
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="mb-3">
                                {!! Form::label('email', 'Email', ['class' => 'form-label']) !!}
                                <span class="text-danger">*</span>
                                {!! Form::email('email', old('email'), [
                                'class' => 'form-control' . ($errors->has('email') ? ' is-invalid' : ''),
                                'id' => 'email',
                                'required' => true,
                                'placeholder' => 'Enter email address'
                                ]) !!}
                                @error('email')
                                <div class="invalid-feedback">{{ $message }}</div>
                                @enderror
                            </div>
                        </div>
                    </div>

                    <div class="row">
                        <div class="col-md-6">
                            <div class="mb-3">
                                {!! Form::label('phone', 'Phone', ['class' => 'form-label']) !!}
                                <span class="text-danger">*</span>
                                {!! Form::tel('phone', old('phone'), [
                                'class' => 'form-control' . ($errors->has('phone') ? ' is-invalid' : ''),
                                'id' => 'phone',
                                'required' => true,
                                'placeholder' => 'Enter phone number'
                                ]) !!}
                                @error('phone')
                                <div class="invalid-feedback">{{ $message }}</div>
                                @enderror
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div class="mb-3">
                                {!! Form::label('date_of_birth', 'Date of Birth', ['class' => 'form-label']) !!}
                                <span class="text-danger">*</span>
                                {!! Form::date('date_of_birth', old('date_of_birth'), [
                                'class' => 'form-control' . ($errors->has('date_of_birth') ? ' is-invalid' : ''),
                                'id' => 'date_of_birth',
                                'required' => true,
                                'max' => date('Y-m-d')
                                ]) !!}
                                @error('date_of_birth')
                                <div class="invalid-feedback">{{ $message }}</div>
                                @enderror
                            </div>
                        </div>
                    </div>

                    <div class="mb-3">
                        {!! Form::label('nationality', 'Nationality', ['class' => 'form-label']) !!}
                        <span class="text-danger">*</span>
                        {!! Form::select('nationality',
                        ['' => 'Select Nationality'] + collect($nationalities)->pluck('country', 'id')->toArray(),
                        old('nationality'),
                        [
                        'class' => 'form-select' . ($errors->has('nationality') ? ' is-invalid' : ''),
                        'id' => 'nationality',
                        'required' => true
                        ]
                        ) !!}
                        @error('nationality')
                        <div class="invalid-feedback">{{ $message }}</div>
                        @enderror
                    </div>

                    <!-- Family List Section -->
                    <div class="mb-4">
                        <div class="d-flex justify-content-between align-items-center mb-3">
                            <h5>Family List</h5>
                            <button type="button" class="btn btn-sm btn-outline-primary" id="addFamilyBtn">
                                <i class="fas fa-plus"></i> Add Family Member
                            </button>
                        </div>

                        <div id="familyList" data-family-count="{{ old('family_list') ? count(old('family_list')) : 0 }}">
                            @if(old('family_list'))
                            @foreach(old('family_list') as $index => $family)
                            <div class="family-item border p-3 mb-3 rounded">
                                <div class="d-flex justify-content-between align-items-center mb-2">
                                    <h6 class="mb-0">Family Member {{ $index + 1 }}</h6>
                                    <button type="button" class="btn btn-sm btn-outline-danger remove-family">
                                        <i class="fas fa-times"></i>
                                    </button>
                                </div>
                                <div class="row">
                                    <div class="col-md-4">
                                        <div class="mb-2">
                                            {!! Form::label("family_list[{$index}][name]", 'Name', ['class' => 'form-label']) !!}
                                            {!! Form::text("family_list[{$index}][name]", $family['name'] ?? '', [
                                            'class' => 'form-control',
                                            'placeholder' => 'Enter family member name'
                                            ]) !!}
                                        </div>
                                    </div>
                                    <div class="col-md-4">
                                        <div class="mb-2">
                                            {!! Form::label("family_list[{$index}][date_of_birth]", 'Date of Birth', ['class' => 'form-label']) !!}
                                            {!! Form::date("family_list[{$index}][date_of_birth]", $family['date_of_birth'] ?? '', [
                                            'class' => 'form-control',
                                            'max' => date('Y-m-d')
                                            ]) !!}
                                        </div>
                                    </div>
                                    <div class="col-md-4">
                                        <div class="mb-2">
                                            {!! Form::label("family_list[{$index}][relation]", 'Relation', ['class' => 'form-label']) !!}
                                            {!! Form::select("family_list[{$index}][relation]", [
                                            '' => 'Select Relation',
                                            'spouse' => 'Spouse',
                                            'child' => 'Child',
                                            'parent' => 'Parent',
                                            'sibling' => 'Sibling',
                                            'other' => 'Other'
                                            ], $family['relation'] ?? '', [
                                            'class' => 'form-select'
                                            ]) !!}
                                        </div>
                                    </div>
                                </div>
                            </div>
                            @endforeach
                            @endif
                        </div>
                    </div>

                    <div class="d-flex justify-content-between">
                        {!! Form::button('<i class="fas fa-arrow-left"></i> Back', [
                        'type' => 'button',
                        'class' => 'btn btn-secondary',
                        'onclick' => 'history.back()'
                        ]) !!}

                        {!! Form::button('<i class="fas fa-save"></i> Save Customer', [
                        'type' => 'submit',
                        'class' => 'btn btn-primary'
                        ]) !!}
                    </div>
                    {!! Form::close() !!}
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    document.addEventListener('DOMContentLoaded', function() {
        // Mengambil family count dari data attribute
        let familyIndex = parseInt(document.getElementById('familyList').getAttribute('data-family-count')) || 0;

        document.getElementById('addFamilyBtn').addEventListener('click', function() {
            const familyList = document.getElementById('familyList');
            const familyItem = document.createElement('div');
            familyItem.className = 'family-item border p-3 mb-3 rounded';
            familyItem.innerHTML = `
            <div class="d-flex justify-content-between align-items-center mb-2">
                <h6 class="mb-0">Family Member ${familyIndex + 1}</h6>
                <button type="button" class="btn btn-sm btn-outline-danger remove-family">
                    <i class="fas fa-times"></i>
                </button>
            </div>
            <div class="row">
                <div class="col-md-4">
                    <div class="mb-2">
                        <label class="form-label">Name</label>
                        <input type="text" class="form-control" name="family_list[${familyIndex}][name]" placeholder="Enter family member name">
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="mb-2">
                        <label class="form-label">Date of Birth</label>
                        <input type="date" class="form-control" name="family_list[${familyIndex}][date_of_birth]" max="${new Date().toISOString().split('T')[0]}">
                    </div>
                </div>
                <div class="col-md-4">
                    <div class="mb-2">
                        <label class="form-label">Relation</label>
                        <select class="form-select" name="family_list[${familyIndex}][relation]">
                            <option value="">Select Relation</option>
                            <option value="spouse">Spouse</option>
                            <option value="child">Child</option>
                            <option value="parent">Parent</option>
                            <option value="sibling">Sibling</option>
                            <option value="other">Other</option>
                        </select>
                    </div>
                </div>
            </div>
        `;

            familyList.appendChild(familyItem);
            familyIndex++;
        });

        // Remove family member
        document.addEventListener('click', function(e) {
            if (e.target.classList.contains('remove-family') || e.target.closest('.remove-family')) {
                const familyItem = e.target.closest('.family-item');
                familyItem.remove();

                // Update family member numbers
                const familyItems = document.querySelectorAll('.family-item');
                familyItems.forEach((item, index) => {
                    const header = item.querySelector('h6');
                    header.textContent = `Family Member ${index + 1}`;
                });
            }
        });

        // Form validation
        document.getElementById('customerForm').addEventListener('submit', function(e) {
            const name = document.getElementById('name').value.trim();
            const email = document.getElementById('email').value.trim();
            const phone = document.getElementById('phone').value.trim();
            const dateOfBirth = document.getElementById('date_of_birth').value;
            const nationality = document.getElementById('nationality').value;

            if (!name || !email || !phone || !dateOfBirth || !nationality) {
                e.preventDefault();
                alert('Please fill in all required fields.');
                return false;
            }

            // Email validation
            const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            if (!emailRegex.test(email)) {
                e.preventDefault();
                alert('Please enter a valid email address.');
                return false;
            }

            // Phone validation
            const phoneRegex = /^[\d\s\-\+\(\)]+$/;
            if (!phoneRegex.test(phone)) {
                e.preventDefault();
                alert('Please enter a valid phone number.');
                return false;
            }

            // Date validation
            const today = new Date();
            const selectedDate = new Date(dateOfBirth);
            if (selectedDate >= today) {
                e.preventDefault();
                alert('Date of birth must be in the past.');
                return false;
            }

            // Show loading state
            const submitBtn = document.querySelector('button[type="submit"]');
            submitBtn.disabled = true;
            submitBtn.innerHTML = '<i class="fas fa-spinner fa-spin"></i> Saving...';
        });
    });
</script>
@endsection