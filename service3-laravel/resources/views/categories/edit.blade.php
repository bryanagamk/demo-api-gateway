@extends('layouts.app')

@section('content')
<div class="row">
    <div class="col-md-6 offset-md-3">
        <h2>Edit Category</h2>

        @if ($errors->any())
        <div class="alert alert-danger">
            <ul>
                @foreach ($errors->all() as $err)
                <li>{{ $err }}</li>
                @endforeach
            </ul>
        </div>
        @endif

        <form action="{{ route('categories.update', $category->id) }}" method="POST">
            @csrf
            @method('PUT')
            <div class="mb-3">
                <label>Name</label>
                <input type="text" name="name" class="form-control"
                    value="{{ old('name', $category->name) }}" required>
            </div>
            <div class="mb-3">
                <label>Description</label>
                <textarea name="description" class="form-control">{{ old('description', $category->description) }}</textarea>
            </div>
            <button type="submit" class="btn btn-primary">Update</button>
            <a href="{{ route('categories.index') }}" class="btn btn-secondary">Back</a>
        </form>
    </div>
</div>
@endsection