@extends('layouts.app')

@section('content')
<div class="row">
    <div class="col-md-8 offset-md-2">
        <h2>Category Detail</h2>

        <table class="table table-bordered">
            <tr>
                <th>ID</th>
                <td>{{ $category->id }}</td>
            </tr>
            <tr>
                <th>Name</th>
                <td>{{ $category->name }}</td>
            </tr>
            <tr>
                <th>Description</th>
                <td>{{ $category->description }}</td>
            </tr>
        </table>

        <a href="{{ route('categories.index') }}" class="btn btn-secondary">Back</a>
    </div>
</div>
@endsection